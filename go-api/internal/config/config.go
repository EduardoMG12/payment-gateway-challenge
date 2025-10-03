package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	AmqpURI     string
	RedisURI    string
}

func LoadConfig() *Config {

	env := os.Getenv("APP_ENV")

	fmt.Println("Running in " + env + " mode")

	dbURL := dbUrlParser().DatabaseURL
	amqpURI := rabbitMQURIParser().AmqpURI
	redisURI := redisUriParser().RedisURI

	return &Config{
		DatabaseURL: dbURL,
		AmqpURI:     amqpURI,
		RedisURI:    redisURI,
	}
}
