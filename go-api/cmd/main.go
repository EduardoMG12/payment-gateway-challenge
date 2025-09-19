package main

import (
	"fmt"
	"log"
	"net/http"
	"payment-gateway/go-api/internal/config"
	"payment-gateway/go-api/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DatabaseURL)

	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		log.Fatalln(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Server up\nPORT:8080\nhttp://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
