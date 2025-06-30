package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClientWrapper wraps redis.Client to implement RedisInterface
type RedisClientWrapper struct {
	client *redis.Client
}

func NewRedisClientWrapper(client *redis.Client) *RedisClientWrapper {
	return &RedisClientWrapper{client: client}
}

func (w *RedisClientWrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return w.client.Set(ctx, key, value, expiration).Err()
}

func (w *RedisClientWrapper) Get(ctx context.Context, key string) (string, error) {
	return w.client.Get(ctx, key).Result()
}

func (w *RedisClientWrapper) Del(ctx context.Context, keys ...string) error {
	return w.client.Del(ctx, keys...).Err()
}
