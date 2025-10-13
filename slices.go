package predis

import (
	"context"
	"encoding/json"
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

func LRange[T any](ctx context.Context, key string, start, stop int64) ([]T, error) {
	vals, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	items := make([]T, 0, len(vals))
	for _, v := range vals {
		var item T
		if err := json.Unmarshal([]byte(v), &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func LAll[T any](ctx context.Context, key string) ([]T, error) {
	return LRange[T](ctx, key, 0, -1)
}