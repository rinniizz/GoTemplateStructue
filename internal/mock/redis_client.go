package mock

import (
	"context"
	"errors"
	"time"
)

// MockRedisClient จำลอง Redis client
type MockRedisClient struct {
	data map[string]string
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{
		data: make(map[string]string),
	}
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if str, ok := value.(string); ok {
		m.data[key] = str
	}
	return nil
}

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	if value, exists := m.data[key]; exists {
		return value, nil
	}
	return "", errors.New("key not found")
}

func (m *MockRedisClient) Del(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		delete(m.data, key)
	}
	return nil
}
