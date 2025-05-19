package main

import (
	"log"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api"
	"github.com/mitra-gh/CarBid/internal/data/cache"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Redis
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	// Initialize Server
	api.InitServer(cfg)
}