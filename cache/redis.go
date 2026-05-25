package cache

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
}
func ClearTaskCache() {
	keys, err := RedisClient.Keys(Ctx, "tasks:*").Result()

	if err != nil {
		return
	}

	if len(keys) > 0 {
		RedisClient.Del(Ctx, keys...)
	}
}