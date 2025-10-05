package predis

import (
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Setup(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
