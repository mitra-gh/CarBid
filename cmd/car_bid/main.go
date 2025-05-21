package main

import (
	"log"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api"
	"github.com/mitra-gh/CarBid/internal/data/cache"
	"github.com/mitra-gh/CarBid/internal/data/db"
	"github.com/mitra-gh/CarBid/pkg/logging"
)



func main() {

	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.StartUp, "Application started", nil)

	// Initialize Redis
	err = cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Cache, logging.StartUp, err.Error(),nil)
	}
	defer cache.CloseRedis()

	// Initialize Database
	err = db.InitDB(cfg)
	if err != nil {
		logger.Fatal(logging.Database, logging.StartUp, err.Error(),nil)
	}
	defer db.CloseDB()

	// Initialize Server
	api.InitServer(cfg)
}