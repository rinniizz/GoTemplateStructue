package database

import (
	"context"
	"fmt"

	"go-template-structure/internal/config"
	"go-template-structure/pkg/logger"

	"github.com/go-redis/redis/v8"
)

// NewRedisConnection creates a new Redis connection
func NewRedisConnection(cfg config.RedisConfig) (*RedisClientWrapper, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Test the connection
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	logger.Info("Successfully connected to Redis")
	return NewRedisClientWrapper(rdb), nil
}
