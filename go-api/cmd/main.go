package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/config"
	"payment-gateway/go-api/internal/database"
	"payment-gateway/go-api/internal/router"
	"payment-gateway/go-api/internal/transaction"
	"payment-gateway/go-api/internal/utils"

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

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		log.Fatalln(err)
	}
	defer db.Close()

	mqClient, err := utils.NewRabbitMQClient(cfg.AmqpURI)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	accountModule := account.NewModule(db)
	cardModule := *card.NewModule(db, accountModule.Service)
	transactionModule := transaction.NewModule(db, accountModule.Service, mqClient, cardModule.Service)

	r := router.NewRouter(accountModule.Handler, cardModule.Handler, transactionModule.Handler)
	r.RegisterRoutes()

	fmt.Println("Server running 🚀🚀🚀   PORT:8080")
	fmt.Println("go-api: http://localhost:" + os.Getenv("API_PORT"))
	fmt.Printf("API Swagger doc up: http://localhost:" + os.Getenv("API_PORT") + "/swagger/index.html\n")

	log.Fatal(http.ListenAndServe(":8080", r.MuxRouter()))
}
