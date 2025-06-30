package interfaces

import (
	"context"
	"time"
)

// RedisInterface defines methods for Redis operations
type RedisInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
}
