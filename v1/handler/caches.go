package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	balance "github.com/m3o/services/balance/proto"
	publicapi "github.com/m3o/services/publicapi/proto"
	usage "github.com/m3o/services/usage/proto"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/logger"
)

// publicapiCache caches the apis so we don't hit the publicapi for every call
type publicapiCache struct {
	sync.RWMutex
	apis map[string]*publicapi.PublicAPI
	papi publicapi.PublicapiService
}

func (p *publicapiCache) getPrice(api, endpoint string) int64 {
	p.RLock()
	defer p.RUnlock()
	return p.apis[strings.ToLower(api)].Pricing[strings.ToLower(fmt.Sprintf("%s.%s", api, endpoint))]
}

func (p *publicapiCache) getQuota(api, endpoint string) int64 {
	p.RLock()
	defer p.RUnlock()
	a := p.apis[strings.ToLower(api)]
	if a == nil {
		return 0
	}
	return a.Quotas[strings.ToLower(fmt.Sprintf("%s.%s", api, endpoint))]
}

func (p *publicapiCache) list() ([]*publicapi.PublicAPI, error) {
	rsp, err := p.papi.List(context.Background(), &publicapi.ListRequest{}, client.WithAuthToken())
	return rsp.Apis, err
}

func (p *publicapiCache) init() error {
	// load up the cache and periodically refresh
	load := func() error {
		rsp, err := p.papi.List(context.Background(), &publicapi.ListRequest{}, client.WithAuthToken())
		if err != nil {
			logger.Errorf("Failed to load publicapi pricing %s", err)
			return err
		}
		newMap := map[string]*publicapi.PublicAPI{}
		for _, api := range rsp.Apis {
			// normalise the maps for easier search
			for name, price := range api.Pricing {
				api.Pricing[strings.ToLower(name)] = price
			}
			for name, quota := range api.Quotas {
				api.Quotas[strings.ToLower(name)] = quota
			}
			newMap[strings.ToLower(api.Name)] = api
		}
		p.Lock()
		p.apis = newMap
		p.Unlock()

		return nil
	}
	if err := load(); err != nil {
		return err
	}
	go func() {
		for {
			time.Sleep(2 * time.Minute)
			load()
		}
	}()
	return nil

}

// balanceCache caches the calls to balance
type balanceCache struct {
	balsvc balance.BalanceService
}

func (b *balanceCache) getBalance(ctx context.Context, userID string) (int64, error) {
	// TODO caching
	rsp, err := b.balsvc.Current(ctx, &balance.CurrentRequest{
		CustomerId: userID,
	}, client.WithAuthToken())
	if err != nil {
		return 0, err
	}
	return rsp.CurrentBalance, nil
}

// expiringLRUCache caches the API key records for faster retrieval rather than hitting the DB for every call
type expiringLRUCache struct {
	redisClient *redis.Client
	ttl         time.Duration
}

func (c *expiringLRUCache) Add(ctx context.Context, key string, value interface{}) error {
	val, _ := json.Marshal(value)
	return c.redisClient.Set(ctx, fmt.Sprintf("%s:%s", keyCachePrefix, key), val, c.ttl).Err()
}

func (c *expiringLRUCache) Remove(ctx context.Context, key string) error {
	return c.redisClient.Del(ctx, fmt.Sprintf("%s:%s", keyCachePrefix, key)).Err()
}

func (c *expiringLRUCache) GetAPIKeyRecord(ctx context.Context, key string) (*apiKeyRecord, error) {
	b, err := c.redisClient.Get(ctx, fmt.Sprintf("%s:%s", keyCachePrefix, key)).Bytes()
	if err != nil {
		return nil, err
	}
	var keyRec apiKeyRecord
	if err := json.Unmarshal(b, &keyRec); err != nil {
		return nil, err
	}
	return &keyRec, nil
}

type usageCache struct {
	usagesvc usage.UsageService
}

func (u *usageCache) getMonthlyUsageTotal(ctx context.Context, userID string) (map[string]int64, error) {
	rsp, err := u.usagesvc.ReadMonthlyTotal(ctx, &usage.ReadMonthlyTotalRequest{
		CustomerId: userID,
		Detail:     true,
	}, client.WithAuthToken())
	if err != nil {
		return nil, err
	}
	return rsp.EndpointRequests, nil
}
