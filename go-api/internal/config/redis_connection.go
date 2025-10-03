package config

import (
	"fmt"
	"os"
)

type RedisUri struct {
	RedisURI string
}

func redisUriParser() *RedisUri {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	fmt.Print("\nredis uri:", redisAddr)

	return &RedisUri{
		RedisURI: redisAddr,
	}
}
