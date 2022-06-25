package redis

import (
	"errors"

	"github.com/go-redis/redis/v7"
)

var RedisConn *redis.Client

func Connect() error {
	RedisConn = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	_, err := RedisConn.Ping().Result()
	if err != nil {
		return errors.New("Cannot Connect")
	}

	return nil
}
