package main

import (
	"fmt"
	"log"
	"net/http"

	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/config"
	"payment-gateway/go-api/internal/database"
	"payment-gateway/go-api/internal/repository"
	"payment-gateway/go-api/internal/router"

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

	accountRepo := repository.NewAccountRepository(db)
	accountService := account.NewAccountService(accountRepo)
	accountHandler := account.NewAccountHandler(accountService)

	cardRepo := repository.NewCardRepository(db)
	cardService := card.NewCardService(cardRepo, accountService)
	cardHandler := card.NewCardHandler(cardService)

	r := router.NewRouter(accountHandler, cardHandler)
	r.RegisterRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server up\nPORT:8080\nhttp://localhost:8080")
	fmt.Printf("API Swagger doc up: http://localhost:8080/swagger")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
