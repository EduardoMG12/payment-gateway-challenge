package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	AmqpURI     string
}

func LoadConfig() *Config {

	env := os.Getenv("APP_ENV")

	fmt.Println("Running in " + env + " mode")

	dbURL := db_connection().DatabaseURL
	amqpURI := rabbitmq_connection().AmqpURI

	return &Config{
		DatabaseURL: dbURL,
		AmqpURI:     amqpURI,
	}
}
