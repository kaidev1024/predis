package predis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func SetStruct(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	return set(ctx, key, data, coldDefaultExpiration)
}

func SetStructWithExpiration(ctx context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	return set(ctx, key, data, expiration)
}

// GetStruct retrieves JSON and unmarshals into a pointer type using generics
func GetStruct[T any](ctx context.Context, key string, dest *T) error {
	val, err := get(ctx, key)
	if err != nil {
		return fmt.Errorf("redis get error: %w", err)
	}
	return json.Unmarshal([]byte(val), dest)
}

func HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.HSet(ctx, key, values...).Result()
}

func HGet(ctx context.Context, key, field string) (string, error) {
	return client.HGet(ctx, key, field).Result()
}

func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return client.HGetAll(ctx, key).Result()
}
