package card

import (
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Module struct {
	Handler *CardHandler
	Service CardService
}

func NewModule(db *sqlx.DB, accountService account.AccountService) *Module {
	repo := repository.NewCardRepository(db)
	service := NewCardService(repo, accountService)
	handler := NewCardHandler(service)

	return &Module{
		Handler: handler,
		Service: service,
	}
}
