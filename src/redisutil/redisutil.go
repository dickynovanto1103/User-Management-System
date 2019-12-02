package redisutil

import (
	"config"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func CreateRedisClient(config config.ConfigRedis) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	pong, err := redisClient.Ping().Result()
	log.Println(pong, err)
}

func Set(key string, value string, duration time.Duration) {
	redisClient.Set(key, value, duration)
}

func Get(key string) (string, error) {
	result, err := redisClient.Get(key).Result()
	return result, err
}
