package main

import (
	"m3o.dev/platform/service"
	"m3o.dev/platform/service/logger"
	"m3o.dev/services/chat/handler"
	pb "m3o.dev/services/chat/proto"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("chat"),
		service.Version("latest"),
	)

	// Register the handler against the server
	pb.RegisterChatHandler(srv.Server(), new(handler.Chat))

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
