package cache

import (
	"time"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *configs.Config) *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
		DialTimeout: cfg.Redis.DialTimeout * time.Second,
		ReadTimeout: cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize: cfg.Redis.PoolSize,
		PoolTimeout: cfg.Redis.PoolTimeout * time.Second,
	})
	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() error {
	return redisClient.Close()
}

