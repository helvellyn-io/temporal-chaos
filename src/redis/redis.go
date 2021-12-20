package redis

import (
	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "2.tcp.ngrok.io:11500",
		Password: "",
		DB:       0,
	})
	return client
}
