package config

import (
	"log"
	"os"
)

type DbUrl struct {
	DatabaseURL string
}

func dbUrlParser() *DbUrl {

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")

	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":5432/" + dbName + "?sslmode=disable"
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not configured.")
	}

	return &DbUrl{
		DatabaseURL: dbURL,
	}

}
