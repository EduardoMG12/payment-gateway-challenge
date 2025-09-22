package router

import (
	"net/http"
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	AccountHandler *account.AccountHandler
	CardHandler    *card.CardHandler
}

func NewRouter(accountHandler *account.AccountHandler, cardHandler *card.CardHandler) *Router {
	return &Router{
		AccountHandler: accountHandler,
		CardHandler:    cardHandler,
	}
}

func (r *Router) RegisterRoutes() {
	http.HandleFunc("/accounts", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.AccountHandler.CreateAccount(w, req)
		case "GET":
			r.AccountHandler.GetAllAccounts(w, req)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cards", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.CardHandler.CreateCard(w, req)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)
}
