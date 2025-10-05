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

func Get(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshalling the struct")
	}
	return set(ctx, key, data, expiration)
}
