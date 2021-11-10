package handler

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	custpb "github.com/m3o/services/customers/proto"
	m3oauth "github.com/m3o/services/pkg/auth"
	publicapipb "github.com/m3o/services/publicapi/proto"
	pb "github.com/m3o/services/usage/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	dbproto "github.com/micro/services/db/proto"
)

const (
	prefixCounter         = "usage-service/counter"
	prefixUsageByCustomer = "usageByCustomer" // customer ID / date
	counterTTL            = 48 * time.Hour
	counterMonthlyTTL     = 40 * 24 * time.Hour
)

type UsageSvc struct {
	sync.RWMutex
	c               *counter
	dbService       dbproto.DbService
	custService     custpb.CustomersService
	papiService     publicapipb.PublicapiService
	rankCache       []*pb.APIRankItem
	globalRankCache []*pb.APIRankUserItem
}

func NewHandler(svc *service.Service, dbService dbproto.DbService) *UsageSvc {
	redisConfig := struct {
		Address  string
		User     string
		Password string
	}{}
	val, err := config.Get("micro.usage.redis")
	if err != nil {
		log.Fatalf("No redis config found %s", err)
	}
	if err := val.Scan(&redisConfig); err != nil {
		log.Fatalf("Error parsing redis config %s", err)
	}
	if len(redisConfig.Password) == 0 || len(redisConfig.User) == 0 || len(redisConfig.Password) == 0 {
		log.Fatalf("Missing redis config %s", err)
	}
	rc := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Username: redisConfig.User,
		Password: redisConfig.Password,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	})
	p := &UsageSvc{
		c:           &counter{redisClient: rc},
		dbService:   dbService,
		rankCache:   []*pb.APIRankItem{},
		custService: custpb.NewCustomersService("customers", svc.Client()),
		papiService: publicapipb.NewPublicapiService("publicapi", svc.Client()),
	}
	p.RankingCron()
	go p.consumeEvents()
	return p
}

func (p *UsageSvc) Read(ctx context.Context, request *pb.ReadRequest, response *pb.ReadResponse) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("usage.Read", "Unauthorized")
	}
	if len(request.CustomerId) == 0 {
		request.CustomerId = acc.ID
	}
	if acc.ID != request.CustomerId {
		_, err := m3oauth.VerifyMicroAdmin(ctx, "usage.Read")
		if err != nil {
			return err
		}
	}

	now := time.Now().UTC().Truncate(24 * time.Hour)
	liveEntries, err := p.c.listForUser(request.CustomerId, now)
	if err != nil {
		log.Errorf("Error retrieving usage %s", err)
		return errors.InternalServerError("usage.Read", "Error retrieving usage")
	}

	response.Usage = map[string]*pb.UsageHistory{}
	// add live data on top of historical
	keyPrefix := fmt.Sprintf("%s/%s/", prefixUsageByCustomer, request.CustomerId)
	recs, err := store.Read(keyPrefix, store.ReadPrefix())
	if err != nil {
		log.Errorf("Error querying historical data %s", err)
		return errors.InternalServerError("usage.Read", "Error retrieving usage")
	}

	addEntryToResponse := func(response *pb.ReadResponse, e listEntry, unixTime int64) {
		// detailed view includes data for individual endpoints
		if !request.Detail && strings.Contains(e.Service, "$") {
			return
		}
		use := response.Usage[e.Service]
		if use == nil {
			use = &pb.UsageHistory{
				ApiName: e.Service,
				Records: []*pb.UsageRecord{},
			}
		}
		use.Records = append(use.Records, &pb.UsageRecord{Date: unixTime, Requests: e.Count})
		response.Usage[e.Service] = use
	}

	// add to slices
	for _, rec := range recs {
		date := strings.TrimPrefix(rec.Key, keyPrefix)
		if len(date) != 8 {
			continue
		}
		dateObj, err := time.Parse("20060102", date)
		if err != nil {
			log.Errorf("Error parsing date obj %s", err)
			return errors.InternalServerError("usage.Read", "Error retrieving usage")
		}
		var de dateEntry
		if err := json.Unmarshal(rec.Value, &de); err != nil {
			log.Errorf("Error parsing date obj %s", err)
			return errors.InternalServerError("usage.Read", "Error retrieving usage")
		}
		for _, e := range de.Entries {
			addEntryToResponse(response, e, dateObj.Unix())
		}
	}
	for _, e := range liveEntries {
		addEntryToResponse(response, e, now.Unix())
	}
	// sort slices
	for _, v := range response.Usage {
		sort.Slice(v.Records, func(i, j int) bool {
			if v.Records[i].Date == v.Records[j].Date {
				return v.Records[i].Requests < v.Records[j].Requests
			}
			return v.Records[i].Date < v.Records[j].Date
		})
	}
	// remove dupe
	for k, v := range response.Usage {
		lenRecs := len(v.Records)
		if lenRecs < 2 {
			continue
		}
		if v.Records[lenRecs-2].Date != v.Records[lenRecs-1].Date {
			continue
		}
		response.Usage[k].Records = append(v.Records[:lenRecs-2], v.Records[lenRecs-1])

	}
	return nil
}

type dateEntry struct {
	Entries []listEntry
}

func (p *UsageSvc) UsageCron() {
	defer func() {
		log.Infof("Usage sweep ended")
	}()
	log.Infof("Performing usage sweep")
	// loop through counters and persist
	ctx := context.Background()
	sc := p.c.redisClient.Scan(ctx, 0, prefixCounter+":*", 0)
	if err := sc.Err(); err != nil {
		log.Errorf("Error running redis scan %s", err)
		return
	}

	toPersist := map[string]map[string][]listEntry{} // userid->date->[]listEntry
	it := sc.Iterator()
	for {
		if !it.Next(ctx) {
			if err := it.Err(); err != nil {
				log.Errorf("Error during iteration %s", err)
			}
			break
		}

		key := it.Val()
		count, err := p.c.redisClient.Get(ctx, key).Int64()
		if err != nil {
			log.Errorf("Error retrieving value %s", err)
			return
		}
		parts := strings.Split(strings.TrimPrefix(key, prefixCounter+":"), ":")
		if len(parts) < 3 {
			log.Errorf("Unexpected number of components in key %s", key)
			continue
		}
		userID := parts[0]
		date := parts[1]
		service := parts[2]
		dates := toPersist[userID]
		if dates == nil {
			dates = map[string][]listEntry{}
			toPersist[userID] = dates
		}
		entries := dates[date]
		if entries == nil {
			entries = []listEntry{}
		}
		entries = append(entries, listEntry{
			Service: service,
			Count:   count,
		})
		dates[date] = entries
	}

	for userID, v := range toPersist {
		for date, entry := range v {
			de := dateEntry{
				Entries: entry,
			}
			b, err := json.Marshal(de)
			if err != nil {
				log.Errorf("Error marshalling entry %s", err)
				return
			}
			store.Write(&store.Record{
				Key:   fmt.Sprintf("%s/%s/%s", prefixUsageByCustomer, userID, date),
				Value: b,
			})
		}

	}

}

func (p *UsageSvc) Sweep(ctx context.Context, request *pb.SweepRequest, response *pb.SweepResponse) error {
	p.UsageCron()
	return nil
}

func (p *UsageSvc) deleteUser(ctx context.Context, userID string) error {
	if err := p.c.deleteUser(ctx, userID); err != nil {
		return err
	}

	recs, err := store.Read(fmt.Sprintf("%s/%s/", prefixUsageByCustomer, userID), store.ReadPrefix())
	if err != nil {
		return err
	}
	for _, rec := range recs {
		if err := store.Delete(rec.Key); err != nil {
			return err
		}
	}
	return nil

}

func (p *UsageSvc) DeleteCustomer(ctx context.Context, request *pb.DeleteCustomerRequest, response *pb.DeleteCustomerResponse) error {
	if _, err := m3oauth.VerifyMicroAdmin(ctx, "usage.DeleteCustomer"); err != nil {
		return err
	}

	if len(request.Id) == 0 {
		return errors.BadRequest("usage.DeleteCustomer", "Error deleting customer")
	}

	if err := p.deleteUser(ctx, request.Id); err != nil {
		log.Errorf("Error deleting customer %s", err)
		return err
	}
	return nil
}

func (p *UsageSvc) SaveEvent(ctx context.Context, request *pb.SaveEventRequest, response *pb.SaveEventResponse) error {
	if request.Event == nil {
		return fmt.Errorf("event not provided")
	}
	if request.Event.Table == "" {
		return fmt.Errorf("table not provided")
	}
	rec := request.Event.Record.AsMap()
	if request.Event.Id == "" {
		request.Event.Id = uuid.New().String()
	}
	rec["id"] = request.Event.Id
	rec["createdAt"] = time.Now().Unix()
	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, request.Event.Record)
	if err != nil {
		return err
	}
	_, err = p.dbService.Create(ctx, &dbproto.CreateRequest{
		Table:  request.Event.Table,
		Record: request.Event.Record,
	})
	return err
}

func (p *UsageSvc) ListEvents(ctx context.Context, request *pb.ListEventsRequest, response *pb.ListEventsResponse) error {
	if request.Table == "" {
		return fmt.Errorf("no table provided")
	}
	resp, err := p.dbService.Read(ctx, &dbproto.ReadRequest{
		Table:   request.Table,
		Query:   "createdAt > 0",
		OrderBy: "createdAt",
		Order:   "desc",
	})
	if err != nil {
		return err
	}
	for _, v := range resp.Records {
		response.Events = append(response.Events, &pb.Event{
			Table:  request.Table,
			Record: v,
		})
	}
	return nil
}

// ListAPIRanks returns the ranking of the APIs based on data up to but not including today.
// We calculate the "popularity" (a score 0-10) based on how many requests an API has compared to the "top" performing API.
// Note: the top performing API is found after removing any APIs with zscore above 1.5 (i.e. outside of 1.5 standard deviations)
func (p *UsageSvc) ListAPIRanks(ctx context.Context, request *pb.ListAPIRanksRequest, response *pb.ListAPIRanksResponse) error {
	p.RLock()
	defer p.RUnlock()
	response.Ranks = p.rankCache
	response.GlobalTopUsers = p.globalRankCache
	return nil
}

func (p *UsageSvc) RankingCron() {
	// construct ranks
	keyPrefix := fmt.Sprintf("%s/%s/", prefixUsageByCustomer, totalID)
	recs, err := store.Read(keyPrefix, store.ReadPrefix())
	if err != nil {
		log.Errorf("Error querying historical data %s", err)
		return
	}

	type entry struct {
		requestCount int64
		apiName      string
	}

	runningTotals := map[string]*entry{}

	thirtyDays := time.Now().Add(-30 * 24 * time.Hour)
	// add to slices
	for _, rec := range recs {
		date := strings.TrimPrefix(rec.Key, keyPrefix)
		dateObj, err := time.Parse("20060102", date)
		if err != nil {
			log.Errorf("Error parsing date obj %s", err)
			return
		}
		if dateObj.Before(thirtyDays) {
			continue
		}
		var de dateEntry
		if err := json.Unmarshal(rec.Value, &de); err != nil {
			log.Errorf("Error parsing date obj %s", err)
			return
		}

		for _, e := range de.Entries {
			if strings.Contains(e.Service, "$") {
				continue
			}
			en := runningTotals[e.Service]
			if en == nil {
				en = &entry{apiName: e.Service}
			}
			en.requestCount += e.Count
			runningTotals[e.Service] = en
		}
	}

	sortSlice := []*entry{}
	reqTotal := int64(0)
	for _, v := range runningTotals {
		sortSlice = append(sortSlice, v)
		reqTotal += v.requestCount
	}
	mean := float64(reqTotal) / float64(len(sortSlice))
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].requestCount > sortSlice[j].requestCount
	})

	sd := float64(0)
	for _, v := range sortSlice {
		sd += math.Pow(float64(v.requestCount)-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(sortSlice)))

	maxReqs := sortSlice[0].requestCount
	// find highest request count with zscore within 3
	for _, v := range sortSlice {
		zscore := (float64(v.requestCount) - mean) / sd
		if zscore <= 1.5 {
			maxReqs = v.requestCount
			break
		}
	}

	retSlice := []*pb.APIRankItem{}

	calcPopularity := func(requestCount int64) int32 {
		ret := int32(math.Round(float64(requestCount) / float64(maxReqs) * 10))
		if ret > 10 {
			return 10
		}
		return ret
	}
	top10s, globalTop10, err := p.calcTop10s()
	if err != nil {
		return
	}

	prsp, err := p.papiService.List(context.Background(), &publicapipb.ListRequest{}, client.WithAuthToken())
	if err != nil {
		log.Errorf("Error retrieving publicapi list %s", err)
		return
	}
	displayNames := map[string]string{}
	for _, api := range prsp.Apis {
		displayNames[api.Name] = api.DisplayName
	}

	for i, v := range sortSlice {
		retSlice = append(retSlice, &pb.APIRankItem{
			ApiName:        v.apiName,
			Position:       int32(i + 1),
			TopUsers:       top10s[v.apiName],
			Popularity:     calcPopularity(v.requestCount),
			ApiDisplayName: displayNames[v.apiName],
		})
	}

	p.Lock()
	p.rankCache = retSlice
	p.globalRankCache = globalTop10
	p.Unlock()
}

func (p *UsageSvc) calcTop10s() (map[string][]*pb.APIRankUserItem, []*pb.APIRankUserItem, error) {
	// construct ranks
	keyPrefix := fmt.Sprintf("%s/", prefixUsageByCustomer)
	recs, err := store.Read(keyPrefix, store.ReadPrefix())
	if err != nil {
		log.Errorf("Error querying historical data %s", err)
		return nil, nil, err
	}

	type userEntry struct {
		id          string
		apiRequests map[string]int64
	}
	users := map[string]*userEntry{} // userID to entry

	type top10User struct {
		id    string
		count int64
	}
	type entry struct {
		top10 []*top10User
	}

	runningTotals := map[string]*entry{}

	thirtyDays := time.Now().Add(-30 * 24 * time.Hour)

	// load up the user counts
	for _, rec := range recs {
		split := strings.Split(rec.Key, "/")
		userID := split[1]
		if userID == totalID {
			// skip the "total" user ID
			continue
		}
		date := strings.TrimPrefix(rec.Key, fmt.Sprintf("%s/%s/", prefixUsageByCustomer, userID))
		if len(date) != 8 {
			// skip anything that isn't a daily total
			continue
		}
		dateObj, err := time.Parse("20060102", date)
		if err != nil {
			log.Errorf("Error parsing date obj %s", err)
			return nil, nil, err
		}
		if dateObj.Before(thirtyDays) {
			continue
		}
		var de dateEntry
		if err := json.Unmarshal(rec.Value, &de); err != nil {
			log.Errorf("Error parsing record %s", err)
			return nil, nil, err
		}
		ue := users[userID]
		if ue == nil {
			ue = &userEntry{
				id:          userID,
				apiRequests: map[string]int64{},
			}
		}

		for _, e := range de.Entries {
			if strings.Contains(e.Service, "$") {
				// skip the endpoint breakdown
				continue
			}
			ue.apiRequests[e.Service] += e.Count
		}
		users[userID] = ue
	}

	globalTopUsers := []*top10User{}
	// calc the top 10s
	for _, user := range users {
		userTotal := int64(0)
		for svc, count := range user.apiRequests {
			userTotal += count
			en := runningTotals[svc]
			userEntry := &top10User{
				id:    user.id,
				count: count,
			}
			if en == nil {
				en = &entry{
					top10: []*top10User{userEntry},
				}
			} else {
				en.top10 = append(en.top10, userEntry)
				sort.Slice(en.top10, func(i, j int) bool {
					return en.top10[i].count > en.top10[j].count
				})
				if len(en.top10) > 10 {
					en.top10 = en.top10[:10]
				}
			}
			runningTotals[svc] = en
		}
		globalTopUsers = append(globalTopUsers, &top10User{
			id:    user.id,
			count: userTotal,
		})
		sort.Slice(globalTopUsers, func(i, j int) bool {
			return globalTopUsers[i].count > globalTopUsers[j].count
		})
		if len(globalTopUsers) > 10 {
			globalTopUsers = globalTopUsers[:10]
		}
	}

	ret := map[string][]*pb.APIRankUserItem{}
	custIDToName := map[string]string{}
	for svc, v := range runningTotals {
		log.Infof("Top 10 %s", svc)
		top := []*pb.APIRankUserItem{}
		for i, s := range v.top10 {
			log.Infof("%+v", s)
			name, err := p.getUserName(context.Background(), s.id, custIDToName)
			if err != nil {
				return nil, nil, err
			}
			top = append(top, &pb.APIRankUserItem{
				UserName: name,
				Position: int32(i + 1),
			})
		}
		ret[svc] = top
	}

	retGlobal := []*pb.APIRankUserItem{}
	log.Infof("Global top 10")
	for i, v := range globalTopUsers {
		log.Infof("%+v", v)
		name, err := p.getUserName(context.Background(), v.id, custIDToName)
		if err != nil {
			return nil, nil, err
		}
		retGlobal = append(retGlobal, &pb.APIRankUserItem{Position: int32(i + 1), UserName: name})
	}
	return ret, retGlobal, nil
}

func (p *UsageSvc) getUserName(ctx context.Context, id string, cache map[string]string) (string, error) {
	name := cache[id]
	if len(name) == 0 {
		rsp, err := p.custService.Read(ctx, &custpb.ReadRequest{Id: id}, client.WithAuthToken())
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				log.Errorf("Error reading customer %s %s", id, err)
				return "", err
			}
			name = "Anonymous"
		} else {
			name = rsp.Customer.Name
			if len(name) == 0 {
				name = rsp.Customer.Meta["generated_name"]
			}
		}
		cache[id] = name
	}
	return name, nil
}

func (p *UsageSvc) ReadMonthlyTotal(ctx context.Context, request *pb.ReadMonthlyTotalRequest, response *pb.ReadMonthlyTotalResponse) error {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("usage.ReadMonthlyTotal", "Unauthorized")
	}
	if len(request.CustomerId) == 0 {
		request.CustomerId = acc.ID
	}
	if acc.ID != request.CustomerId {
		_, err := m3oauth.VerifyMicroAdmin(ctx, "usage.ReadMonthlyTotal")
		if err != nil {
			return err
		}
	}

	count, err := p.c.readMonthly(ctx, request.CustomerId, "total", time.Now())
	if err != nil {
		log.Errorf("Error reading usage %s", err)
		return errors.InternalServerError("usage.ReadMonthlyTotal", "Error reading usage")
	}
	response.Requests = count

	if request.Detail {
		usage, err := p.c.listMonthliesForUser(request.CustomerId, time.Now())
		if err != nil {
			log.Errorf("Error reading usage %s", err)
			return errors.InternalServerError("usage.ReadMonthlyTotal", "Error reading usage")
		}
		ret := map[string]int64{}
		for _, le := range usage {
			ret[le.Service] = le.Count
		}
		response.EndpointRequests = ret
	}
	return nil
}

func (p *UsageSvc) ReadMonthly(ctx context.Context, request *pb.ReadMonthlyRequest, response *pb.ReadMonthlyResponse) error {
	_, err := m3oauth.VerifyMicroAdmin(ctx, "usage.ReadMonthlyTotal")
	if err != nil {
		return err
	}
	t := time.Now()
	response.Requests = map[string]int64{}
	for _, v := range request.Endpoints {
		count, err := p.c.readMonthly(ctx, request.CustomerId, v, t)
		if err != nil {
			log.Errorf("Error reading usage %s", err)
			return errors.InternalServerError("usage.ReadMonthly", "Error reading usage")
		}
		response.Requests[v] = count
	}
	return nil
}
