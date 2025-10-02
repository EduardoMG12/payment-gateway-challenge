package transaction

import (
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/connection"
	"payment-gateway/go-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Module struct {
	Handler *TransactionHandler
}

func NewModule(db *sqlx.DB, accountService account.AccountService, mqClient connection.RabbitMQClient, cardService card.CardService, redis connection.RedisConnection) *Module {
	repo := repository.NewTransactionRepository(db)
	service := NewTransactionService(repo, accountService, mqClient, cardService, redis)
	handler := NewTransactionHandler(service)

	return &Module{
		Handler: handler,
	}
}
