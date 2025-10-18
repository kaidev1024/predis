package predis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RPush[T any](ctx context.Context, key string, items []T) error {
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return err
		}
		if err := client.RPush(ctx, key, data).Err(); err != nil {
			return err
		}
	}
	return nil
}

func LRange[T any](ctx context.Context, key string, start, stop int64) ([]*T, error) {
	data, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, fmt.Errorf("redis lrange error: %w", err)
	}
	n := len(data)
	if n == 0 {
		return nil, redis.Nil
	}
	result := make([]*T, n)
	for i, v := range data {
		item := new(T)
		if err := json.Unmarshal([]byte(v), item); err != nil {
			return nil, err
		}
		result[i] = item
	}
	return result, nil
}

func LAll[T any](ctx context.Context, key string) ([]*T, error) {
	return LRange[T](ctx, key, 0, -1)
}
