package config

import (
	"exchangeapp/global"
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis,got error: %v", err.Error())
	}

	global.RedisDB = RedisClient
}
