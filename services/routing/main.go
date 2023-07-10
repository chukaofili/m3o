package main

import (
	"m3o.dev/services/pkg/tracing"
	"m3o.dev/services/routing/handler"
	pb "m3o.dev/services/routing/proto"

	"m3o.dev/platform/service"
	"m3o.dev/platform/service/config"
	"m3o.dev/platform/service/logger"
	"googlemaps.github.io/maps"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("routing"),
		service.Version("latest"),
	)

	// decide whether to use google or osrm
	c, err := config.Get("routing.provider")
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	// defaults to google
	mode := c.String("")

	switch mode {
	case "google":
		// Setup google maps
		c, err := config.Get("google.apikey")
		if err != nil {
			logger.Fatalf("Error loading config: %v", err)
		}
		apiKey := c.String("")
		if len(apiKey) == 0 {
			logger.Fatalf("Missing required config: google.apikey")
		}
		m, err := maps.NewClient(maps.WithAPIKey(apiKey))
		if err != nil {
			logger.Fatalf("Error configuring google maps client: %v", err)
		}

		// Register handler
		pb.RegisterRoutingHandler(srv.Server(), &handler.Google{m})
	case "osrm", "":
		// todo
		// Setup google maps
		c, err := config.Get("routing.address")
		if err != nil {
			logger.Fatalf("Error loading config: %v", err)
		}
		apiAddr := c.String("http://router.project-osrm.org")
		if len(apiAddr) == 0 {
			logger.Fatalf("Missing required config: routing.address")
		}
		// Register handler
		pb.RegisterRoutingHandler(srv.Server(), &handler.OSRM{Address: apiAddr})
	default:
		logger.Fatalf("%s is an unsupported mode", mode)
	}

	traceCloser := tracing.SetupOpentracing("routing")
	defer traceCloser.Close()
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
