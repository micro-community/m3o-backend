package main

import (
	"github.com/m3o/services/v1api/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/api"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/registry"
	"github.com/micro/micro/v3/service/registry/cache"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("v1"),
		service.Version("latest"),
	)

	srv.Server().Handle(
		srv.Server().NewHandler(
			handler.NewHandler(srv),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "V1.Endpoint",
					Handler: "rpc",
					Method:  []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
					Path:    []string{"^/v1/.*$"},
					Stream:  true,
				}),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "V1.GenerateKey",
					Path:    []string{"/v1/api/keys/generate"},
					Method:  []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
					Handler: "rpc",
				}),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "V1.RevokeKey",
					Path:    []string{"/v1/api/keys/revoke"},
					Method:  []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
					Handler: "rpc",
				}),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "V1.ListKeys",
					Path:    []string{"/v1/api/keys/list"},
					Method:  []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
					Handler: "rpc",
				}),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "V1.ListAPIs",
					Path:    []string{"/v1/api/apis/list"},
					Method:  []string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"},
					Handler: "rpc",
				}),
		))

	// setup caching for registry
	regName := registry.DefaultRegistry.String()
	if regName != "cache" {
		logger.Infof("Setting up cached registry for %s", regName)
		registry.DefaultRegistry = cache.New(registry.DefaultRegistry)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
