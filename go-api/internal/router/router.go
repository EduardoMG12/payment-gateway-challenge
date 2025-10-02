package router

import (
	"net/http"
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/transaction"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	AccountHandler     *account.AccountHandler
	CardHandler        *card.CardHandler
	TransactionHandler *transaction.TransactionHandler
	muxRouter          *mux.Router
}

func (r *Router) MuxRouter() http.Handler {
	return r.muxRouter
}

func NewRouter(accountHandler *account.AccountHandler, cardHandler *card.CardHandler, transactionHandler *transaction.TransactionHandler) *Router {
	return &Router{
		AccountHandler:     accountHandler,
		CardHandler:        cardHandler,
		TransactionHandler: transactionHandler,
		muxRouter:          mux.NewRouter(),
	}
}

func (r *Router) RegisterRoutes() {
	r.muxRouter.HandleFunc("/accounts", r.AccountHandler.CreateAccount).Methods("POST")
	r.muxRouter.HandleFunc("/accounts", r.AccountHandler.GetAllAccounts).Methods("GET")

	r.muxRouter.HandleFunc("/cards", r.CardHandler.CreateCard).Methods("POST")
	r.muxRouter.HandleFunc("/cards/{accountId}", r.CardHandler.GetAllCardsByAccountId).Methods("GET")

	r.muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.muxRouter.HandleFunc("/transactions", r.TransactionHandler.CreateTransaction).Methods("POST")
	r.muxRouter.HandleFunc("/transactions/{accountId}", r.TransactionHandler.GetAllTransactionByAccountIdTestOrderDate).Methods("GET")
	r.muxRouter.HandleFunc("/accounts/{accountId}/balance", r.TransactionHandler.GetBalanceByAccountId).Methods("GET")
}
