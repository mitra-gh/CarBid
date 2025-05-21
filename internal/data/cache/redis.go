package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *configs.Config) error {
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

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() error {
	return redisClient.Close()
}

func Set[T any](key string, value T, expiration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return redisClient.Set(ctx, key, v, expiration*time.Second).Err()
}

func Get[T any](key string) (T, error) {
	ctx := context.Background()
	var dest = *new(T)
	v,err := redisClient.Get(ctx,key).Result()
	if err != nil {
		return dest,err
	}
	err = json.Unmarshal([]byte(v),&dest)
	if err != nil {
		return dest,err
	}
	return dest,nil
}