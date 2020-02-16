package redis

import (
	"log"
	"time"

	"github.com/dickynovanto1103/User-Management-System/internal/service/config"

	"github.com/go-redis/redis"
)

type RedisImpl struct {
	redisClient *redis.Client
}

func CreateRedisClient(config config.ConfigRedis) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	pong, err := redisClient.Ping().Result()
	log.Println(pong, err)
	if err != nil {
		return nil
	}

	return redisClient
}

func CreateRedisWrapper(redisClient *redis.Client) *RedisImpl {
	return &RedisImpl{
		redisClient: redisClient,
	}
}

func (r *RedisImpl) Set(key string, value string, duration time.Duration) {
	r.redisClient.Set(key, value, duration)
}

func (r *RedisImpl) Get(key string) (string, error) {
	result, err := r.redisClient.Get(key).Result()
	return result, err
}
