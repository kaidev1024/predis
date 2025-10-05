package predis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx context.Context

func Setup(addr string, ctx0 context.Context) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", addr),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx = ctx0
}
