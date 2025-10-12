package predis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func SetStruct(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	return set(ctx, key, data, defaultExpiration)
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
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("redis get error: %w", err)
	}
	return json.Unmarshal([]byte(val), dest)
}
