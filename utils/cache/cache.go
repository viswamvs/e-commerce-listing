package cache

import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func InitializeCache() *redis.Client {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RedisClient
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
