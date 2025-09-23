package transaction

import (
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/repository"
	"payment-gateway/go-api/internal/utils"

	"github.com/jmoiron/sqlx"
)

type Module struct {
	Handler *TransactionHandler
}

func NewModule(db *sqlx.DB, accountService account.AccountService, mqClient utils.RabbitMQClient, cardService card.CardService) *Module {
	repo := repository.NewTransactionRepository(db)
	service := NewTransactionService(repo, accountService, mqClient, cardService)
	handler := NewTransactionHandler(service)

	return &Module{
		Handler: handler,
	}
}
