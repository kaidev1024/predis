package predis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var defaultExpiration time.Duration = 24 * time.Hour

func setUpClient(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetUp(addr string) {
	setUpClient(addr)
}

func SetUpWithExpiration(addr string, expiration time.Duration) {
	defaultExpiration = expiration
	setUpClient(addr)
}
