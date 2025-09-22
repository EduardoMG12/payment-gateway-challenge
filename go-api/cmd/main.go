package main

import (
	"fmt"
	"log"
	"net/http"
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/config"
	"payment-gateway/go-api/internal/database"
	"payment-gateway/go-api/internal/repository"

	httpSwagger "github.com/swaggo/http-swagger"

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

	http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			accountHandler.CreateAccount(w, r)
		case "GET":
			accountHandler.GetAllAccounts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server up\nPORT:8080\nhttp://localhost:8080")
	fmt.Printf("API Swagger doc up: http://localhost:8080/swagger")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
