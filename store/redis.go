package store

import (
	"github.com/go-redis/redis"
	"violin-home.cn/retail/config"
)

var ClientRedis *redis.Client

func NewRedisClient() {

	ClientRedis = redis.NewClient(&redis.Options{
		Network:  config.Conf.RC.Network,
		Addr:     config.Conf.RC.Addr,
		Password: config.Conf.RC.Password,
		DB:       config.Conf.RC.DB,
	})

}
