package predis

import "github.com/redis/go-redis/v9"

func IsNil(err error) bool {
	return err == redis.Nil
}
