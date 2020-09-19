package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/service/config"
	mstore "github.com/micro/micro/v3/service/store"

	"github.com/google/uuid"
	alert "github.com/m3o/services/alert/proto/alert"
	"github.com/slack-go/slack"
)

const (
	storePrefixEvents = "events/"
)

type Alert struct {
	slackClient  *slack.Client
	gaPropertyID string
	slackEnabled bool
}

type event struct {
	ID       string            `json:"id"`
	UserID   string            `json:"userID"`
	Category string            `json:"category"`
	Action   string            `json:"action"`
	Label    string            `json:"label"`
	Value    uint64            `json:"value"`
	Metadata map[string]string `json:"metadata"`
}

func NewAlert(store store.Store) *Alert {
	slackToken := config.Get("micro", "alert", "slack_token").String("")
	slackEnabled := config.Get("micro", "alert", "slack_enabled").Bool(true)
	gaPropertyID := config.Get("micro", "alert", "ga_property_id").String("")
	if slackEnabled && len(slackToken) == 0 {
		log.Errorf("Slack token missing")
	}
	if len(gaPropertyID) == 0 {
		log.Errorf("Google Analytics key (property ID) is missing")
	}
	log.Infof("Slack enabled: %v", slackEnabled)

	return &Alert{
		slackClient:  slack.New(slackToken),
		gaPropertyID: gaPropertyID,
		slackEnabled: slackEnabled,
	}
}

// ReportEvent ingests events and sends alerts if needed
func (e *Alert) ReportEvent(ctx context.Context, req *alert.ReportEventRequest, rsp *alert.ReportEventResponse) error {
	if req.Event == nil {
		return errors.New("event can't be empty")
	}
	ev := &event{
		ID:       uuid.New().String(),
		Category: req.Event.Category,
		Action:   req.Event.Action,
		Label:    req.Event.Label,
		Value:    req.Event.Value,
		UserID:   req.Event.UserID,
	}
	// ignoring the error intentionally here so we still sends alerts
	// even if persistence is failing
	err := e.saveEvent(ev)
	if err != nil {
		log.Warnf("Error saving event: %v", err)
	}
	err = e.sendToGA(ev)
	if err != nil {
		log.Warnf("Error sending event to google analytics: %v", err)
	}
	if e.slackEnabled {
		jsond, err := json.MarshalIndent(req.Event, "", "   ")
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("Event received:\n```\n%v\n```", string(jsond))
		_, _, _, err = e.slackClient.SendMessage("errors", slack.MsgOptionUsername("Alert Service"), slack.MsgOptionText(msg, false))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Alert) sendToGA(td *event) error {
	if e.gaPropertyID == "" {
		return errors.New("analytics: GA_TRACKING_ID environment variable is missing")
	}
	if td.Category == "" || td.Action == "" {
		return errors.New("analytics: category and action are required")
	}

	cid := td.UserID
	if len(cid) == 0 {
		// GA does not seem to accept events without user id so we generate a UUID
		cid = uuid.New().String()
	}
	v := url.Values{
		"v":   {"1"},
		"tid": {e.gaPropertyID},
		// Anonymously identifies a particular user. See the parameter guide for
		// details:
		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
		//
		// Depending on your application, this might want to be associated with the
		// user in a cookie.
		"cid": {cid},
		"t":   {"event"},
		"ec":  {td.Category},
		"ea":  {td.Action},
		"ua":  {"cli"},
	}

	if td.Label != "" {
		v.Set("el", td.Label)
	}

	v.Set("ev", fmt.Sprintf("%d", td.Value))

	// NOTE: Google Analytics returns a 200, even if the request is malformed.
	_, err := http.PostForm("https://www.google-analytics.com/collect", v)
	return err
}

func (e *Alert) saveEvent(ev *event) error {
	bytes, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	return mstore.Write(&store.Record{
		Key:   storePrefixEvents + ev.ID,
		Value: bytes})
}
