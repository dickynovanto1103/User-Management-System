package redis

import (
	"log"
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/service/config"

	"github.com/go-redis/redis"
)

type Redis struct {

}

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
