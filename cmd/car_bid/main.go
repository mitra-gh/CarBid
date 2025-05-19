package main

import (
	"log"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api"
	"github.com/mitra-gh/CarBid/internal/data/cache"
	"github.com/mitra-gh/CarBid/internal/data/db"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Redis
	err = cache.InitRedis(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}
	defer cache.CloseRedis()



	// Initialize Database
	err = db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()



	// Initialize Server
	api.InitServer(cfg)
}