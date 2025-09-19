package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@localhost:5432/" + dbName + "?sslmode=disable"
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not configured.")
	}

	return &Config{
		DatabaseURL: dbURL,
	}
}
