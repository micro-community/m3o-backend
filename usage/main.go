package main

import (
	"github.com/m3o/services/usage/handler"

	nsproto "github.com/m3o/services/namespaces/proto"
	pb "github.com/micro/micro/v3/proto/auth"
	rproto "github.com/micro/micro/v3/proto/runtime"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("usage"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(handler.NewUsage(
		nsproto.NewNamespacesService("namespaces", srv.Client()),
		pb.NewAccountsService("auth", srv.Client()),
		rproto.NewRuntimeService("runtime", srv.Client()),
	))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
