package database

import (
	"github.com/go-redis/redis"
	"log"
	"rest_app/app/config"
)

func NewConnectionRedis(config config.AppConfiguration) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Pass,
		DB:       config.Redis.DB,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}

	return redisClient
}
