package main

import (
	"helix/helix/internal/cache"
	"helix/helix/internal/transport/http"
)

func main() {
	// Create cache
	c := cache.Init()

	// Create HTTP transport on port 8080
	t := http.CreateTransport(":8080", c)

	// Start the HTTP server
	t.Init()
}
