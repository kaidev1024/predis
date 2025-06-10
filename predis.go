package predis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func ExampleClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	println("key", val)

	val2, err := redisClient.Get(ctx, "key2").Result()
	if err == redis.Nil {
		println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
