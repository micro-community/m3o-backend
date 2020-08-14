package main

import (
	"github.com/m3o/services/secrets/handler"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
)

func main() {
	srv := service.New(
		service.Name("secrets"),
	)

	srv.Handle(handler.New(srv))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}