package template

var (
	MainSRV = `package main

import (
	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

	"m3o.dev/platform/service"
	"m3o.dev/platform/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("{{lower .Alias}}"),
	)

	// Register handler
	pb.Register{{title .Alias}}Handler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
`
)