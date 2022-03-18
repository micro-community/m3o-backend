package main

import (
	"github.com/m3o/services/admin/handler"
	pb "github.com/m3o/services/admin/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("admin"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAdminHandler(srv.Server(), handler.New(srv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
