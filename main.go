package main

import (
	"cloud.google.com/go"
	pb "go-micro/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go-micro"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGoMicroHandler(srv.Server(), new(handler.GoMicro))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
