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

func LRange[T any](ctx context.Context, key string, start, stop int64, dest []*T) error {
	vals, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return  err
	}


	for _, v := range vals {
		var item T
		if err := json.Unmarshal([]byte(v), &item); err != nil {
			return  err
		}
		dest = append(dest, &item)
	}
	return  nil
}

func LAll[T any](ctx context.Context, key string, dest []*T)  error {
	return LRange(ctx, key, 0, -1, dest)
}