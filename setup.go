package predis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var coldDefaultExpiration time.Duration = time.Hour
var hotDefaultExpiration time.Duration = 24 * time.Hour

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

func SetUpWithExpiration(addr string, coldExpiration, hotExpiration time.Duration) {
	coldDefaultExpiration = coldExpiration
	hotDefaultExpiration = hotExpiration
	setUpClient(addr)
}
