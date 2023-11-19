package store

import (
	"github.com/go-redis/redis"
)

var ClientRedis *redis.Client

func NewRedisClient() {

	ClientRedis = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       3,
	})

}
