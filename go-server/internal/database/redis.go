package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/vitingr/url-shortner/config"
)

func NewRedisClient(secrets *config.Config) (*redis.Client, error) {
	redisDatabase := redis.NewClient(&redis.Options{
		Addr: secrets.RedisAddr,
		Password: "",
		DB: 0,
	})
	
	// Health check: Testa a conexão com o Redis
	ctx := context.Background()
		if err := redisDatabase.Ping(ctx).Err(); err != nil {
		redisDatabase.Close()
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return redisDatabase, nil
}
