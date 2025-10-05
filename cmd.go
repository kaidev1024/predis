package predis

import (
	"encoding/json"
	"fmt"
	"time"
)

func set(key string, value any, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func Set(key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshalling the struct")
	}
	return set(key, data, expiration)
}
