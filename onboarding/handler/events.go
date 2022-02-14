package handler

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	customer "github.com/m3o/services/customers/proto"
	emails "github.com/m3o/services/emails/proto"
	pevents "github.com/m3o/services/pkg/events"
	custpb "github.com/m3o/services/pkg/events/proto/customers"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/config"
	mevents "github.com/micro/micro/v3/service/events"
	logger "github.com/micro/micro/v3/service/logger"
)

type Onboarding struct {
	emailSvc emails.EmailsService
	custSvc  customer.CustomersService
	cfg      conf
}

func NewOnboarding(svc *service.Service) *Onboarding {
	cfg := conf{}
	v, err := config.Get("micro.onboarding")
	if err != nil {
		log.Fatalf("Failed to load config %s", err)
	}
	if err := v.Scan(&cfg); err != nil {
		log.Fatalf("Failed to load config %s", err)
	}
	o := &Onboarding{
		emailSvc: emails.NewEmailsService("emails", svc.Client()),
		custSvc:  customer.NewCustomersService("customers", svc.Client()),
		cfg:      cfg,
	}
	o.consumeEvents()
	return o
}

func (o *Onboarding) consumeEvents() {
	go pevents.ProcessTopic(custpb.Topic, "onboarding", o.processCustomerEvents)
}

func (o *Onboarding) processCustomerEvents(ev mevents.Event) error {
	ctx := context.Background()
	ce := &custpb.Event{}
	if err := json.Unmarshal(ev.Payload, ce); err != nil {
		logger.Errorf("Error unmarshalling customer event: $s", err)
		return nil
	}
	switch ce.Type {
	case custpb.EventType_EventTypeCreated:
		if err := o.processCustomerCreate(ctx, ce); err != nil {
			logger.Errorf("Error processing request event %s", err)
			return err
		}
	default:
		logger.Infof("Skipping event %+v", ce)
	}
	return nil

}

func (o *Onboarding) processCustomerCreate(ctx context.Context, event *custpb.Event) error {
	// Send welcome email
	email := event.Customer.Email
	if len(email) == 0 {
		// look it up
		rsp, err := o.custSvc.Read(ctx, &customer.ReadRequest{Id: event.Customer.Id}, client.WithAuthToken())
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "not found") {
				return nil
			}
			return err
		}
		email = rsp.Customer.Email
	}
	for _, blocked := range o.cfg.EngagementBlockList {
		if blocked == email {
			return nil
		}
	}

	delay := o.cfg.WelcomeDelay
	if len(delay) == 0 {
		delay = "24h"
	}
	dur, _ := time.ParseDuration(delay)

	// TODO add a block list
	if _, err := o.emailSvc.Send(ctx, &emails.SendRequest{
		To:         email,
		TemplateId: o.cfg.Sendgrid.WelcomeTemplateID,
		SendAt:     time.Now().Add(dur).Unix(),
	}, client.WithAuthToken()); err != nil {
		logger.Errorf("Error sending mail %s", err)
		return err
	}

	return nil
}
