package main

import (
	customersproto "github.com/m3o/services/customers/proto"
	inviteproto "github.com/m3o/services/invite/proto"
	nsproto "github.com/m3o/services/namespaces/proto"
	"github.com/m3o/services/signup/handler"
	subproto "github.com/m3o/services/subscriptions/proto"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
	mauth "github.com/micro/micro/v3/service/auth/client"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("signup"),
	)

	// passing in auth because the DefaultAuth is the one used to set up the service
	auth := mauth.NewAuth()

	// Register Handler
	srv.Handle(handler.NewSignup(
		inviteproto.NewInviteService("invite", srv.Client()),
		customersproto.NewCustomersService("customers", srv.Client()),
		nsproto.NewNamespacesService("namespaces", srv.Client()),
		subproto.NewSubscriptionsService("subscriptions", srv.Client()),
		auth,
	))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
