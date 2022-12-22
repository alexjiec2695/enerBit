package database

import (
	"github.com/go-redis/redis"
	"log"
)

func NewConnectionRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "MDNcVb924a",
		DB:       1,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}

	return redisClient
}
