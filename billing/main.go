package main

import (
	"github.com/m3o/services/billing/handler"

	nsproto "github.com/m3o/services/namespaces/proto"
	pproto "github.com/m3o/services/payments/provider/proto"
	subproto "github.com/m3o/services/subscriptions/proto"
	uproto "github.com/m3o/services/usage/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("billing"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.NewBilling(
		nsproto.NewNamespacesService("namespaces", srv.Client()),
		pproto.NewProviderService("payment.stripe", srv.Client()),
		uproto.NewUsageService("usage", srv.Client()),
		subproto.NewSubscriptionsService("subscriptions", srv.Client()),
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
