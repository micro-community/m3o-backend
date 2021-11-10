package handler

import (
	"log"

	balance "github.com/m3o/services/balance/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
)

type Onboarding struct {
	balSvc balance.BalanceService
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
		balSvc: balance.NewBalanceService("balance", svc.Client()),
	}
	return o
}
