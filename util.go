package predis

import (
	"context"
	"time"
)

func set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func get(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}
