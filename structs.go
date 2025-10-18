package predis

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gocql/gocql"
	"github.com/kaidev1024/pugo/pustruct"
	"github.com/redis/go-redis/v9"
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

func HSet[T any](ctx context.Context, key string, obj *T) (int64, error) {
	fields := pustruct.GetFields(obj)
	n := len(fields)
	for i := 1; i < n; i += 2 {
		val := reflect.ValueOf(fields[i])
		typ := val.Type()
		if typ == reflect.TypeOf(gocql.UUID{}) {
			uuidVal := fields[i].(gocql.UUID)
			fields[i] = uuidVal.String()
			continue
		}
		// Handle self-defined types (aliases)
		switch typ.Kind() {
		case reflect.String:
			fields[i] = fmt.Sprintf("%v", val.Interface())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Float32, reflect.Float64, reflect.Bool:
			fields[i] = fmt.Sprintf("%v", val.Interface())
		}
	}
	return client.HSet(ctx, key, fields...).Result()
}

func HGet(ctx context.Context, key, field string) (string, error) {
	return client.HGet(ctx, key, field).Result()
}

func HGetAll[T any](ctx context.Context, key string, target *T) error {
	data, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("redis hgetall error: %w", err)
	}
	if len(data) == 0 {
		return redis.Nil
	}
	return pustruct.UpdateFieldsWithStrings(target, data)
}
