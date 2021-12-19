package redis

import (
	"github.com/go-redis/redis"
)

func RedisClient() (string, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	return pong, err

}
