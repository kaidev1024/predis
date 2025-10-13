package predis

import (
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var coldDefaultExpiration time.Duration = time.Hour
var hotDefaultExpiration time.Duration = 24 * time.Hour

// func setUpClient(addr string) {
// 	client = redis.NewClient(&redis.Options{
// 		Addr:     addr,
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// }

func setUpClient(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("failed to parse redis url: %v", err)
	}
	client = redis.NewClient(opt)
}

func SetUp(url string) {
	setUpClient(url)
}

func SetUpWithExpiration(url string, coldExpiration, hotExpiration time.Duration) {
	coldDefaultExpiration = coldExpiration
	hotDefaultExpiration = hotExpiration
	setUpClient(url)
}
