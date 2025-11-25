package main

import (
	"helix/helix/internal/cache"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}
	defer logger.Sync()

	// Create cache with logger
	cache.Init(logger)
	// keep process alive
	select {}

	// Create HTTP transport on port 8080
	// t := http.CreateTransport(":8080", c)

	// // Start the HTTP server
	// t.Init()
}
