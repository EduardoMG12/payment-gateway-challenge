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
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com o Redis: %w", err)
	}

	fmt.Println("✅ Conectado ao Redis! PORT: localhost:6379")

	return &RedisConnection{
		Client: rdb,
	}, nil
}
