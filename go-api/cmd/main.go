package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/config"
	"payment-gateway/go-api/internal/connection"
	"payment-gateway/go-api/internal/router"
	"payment-gateway/go-api/internal/transaction"

	_ "payment-gateway/go-api/docs"

	_ "github.com/lib/pq"
)

// @title Payment Gateway API
// @version 1.0
// @description This is a main payment-gateway API .
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	cfg := config.LoadConfig()

	db, err := connection.ConnectDatabase(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		log.Fatalln(err)
	}
	defer db.Close()

	mqClient, err := connection.NewRabbitMQClient(cfg.AmqpURI)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	redisConn, err := connection.ConnectRedis(cfg.RedisURI)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	accountModule := account.NewModule(db)
	cardModule := *card.NewModule(db, accountModule.Service)
	transactionModule := transaction.NewModule(db, accountModule.Service, mqClient, cardModule.Service, *redisConn)

	r := router.NewRouter(accountModule.Handler, cardModule.Handler, transactionModule.Handler)
	r.RegisterRoutes()

	handlerWithCors := config.EnableCors(r.MuxRouter())

	fmt.Println("Server running 🚀🚀🚀   PORT:8080")
	fmt.Println("go-api: http://localhost:" + os.Getenv("API_PORT"))
	fmt.Printf("API Swagger doc up: http://localhost:" + os.Getenv("API_PORT") + "/swagger/index.html\n")

	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}
