package router

import (
	"net/http"
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	AccountHandler *account.AccountHandler
	CardHandler    *card.CardHandler
	muxRouter      *mux.Router
}

func (r *Router) MuxRouter() http.Handler {
	return r.muxRouter
}

func NewRouter(accountHandler *account.AccountHandler, cardHandler *card.CardHandler) *Router {
	return &Router{
		AccountHandler: accountHandler,
		CardHandler:    cardHandler,
		muxRouter:      mux.NewRouter(),
	}
}

func (r *Router) RegisterRoutes() {
	r.muxRouter.HandleFunc("/accounts", r.AccountHandler.CreateAccount).Methods("POST")
	r.muxRouter.HandleFunc("/accounts", r.AccountHandler.GetAllAccounts).Methods("GET")

	r.muxRouter.HandleFunc("/cards", r.CardHandler.CreateCard).Methods("POST")
	r.muxRouter.HandleFunc("/cards/{accountId}", r.CardHandler.GetAllCardsByAccountId).Methods("GET")

	r.muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
