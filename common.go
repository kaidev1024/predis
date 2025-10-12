package predis

import (
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func IsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}

func CreateKey(tableName string, cols ...string) string {
	var key string
	for i, col := range cols {
		if i == 0 {
			key = fmt.Sprintf("%s:%s", tableName, col)
		} else {
			key = fmt.Sprintf("%s_%s", key, col)
		}
	}
	return key
}
