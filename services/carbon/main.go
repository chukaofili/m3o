package main

import (
	"m3o.dev/platform/service"
	"m3o.dev/platform/service/logger"
	"m3o.dev/services/carbon/handler"
	pb "m3o.dev/services/carbon/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("carbon"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCarbonHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
