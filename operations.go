package predis

import (
	"context"
	"time"
)

func Del(ctx context.Context, keys ...string) (int64, error) {
	return client.Del(ctx, keys...).Result()
}

func Expire(ctx context.Context, key string) (bool, error) {
	return client.Expire(ctx, key, hotDefaultExpiration).Result()
}

func ExpireWithExpiration(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.Expire(ctx, key, expiration).Result()
}
