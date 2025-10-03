package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Client *redis.Client
}

func ConnectRedis(redisAddr string) (*RedisConnection, error) {
	redisUser := os.Getenv("READER_REDIS_USER")
	redisPassword := os.Getenv("READER_REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: redisUser,
		Password: redisPassword,
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("fail to connect to Redis: %w", err)
	}

	fmt.Println("✅ Connected to Redis! PORT: localhost:6379")

	return &RedisConnection{
		Client: rdb,
	}, nil
}
