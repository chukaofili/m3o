package main

import (
	"m3o.dev/platform/service"
	"m3o.dev/platform/service/logger"
	"m3o.dev/services/postcode/handler"
	pb "m3o.dev/services/postcode/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("postcode"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPostcodeHandler(srv.Server(), new(handler.Postcode))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
