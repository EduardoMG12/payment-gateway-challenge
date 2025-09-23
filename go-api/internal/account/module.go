package account

import (
	"payment-gateway/go-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Module struct {
	Handler *AccountHandler
	Service AccountService
}

func NewModule(db *sqlx.DB) *Module {
	repo := repository.NewAccountRepository(db)
	service := NewAccountService(repo)
	handler := NewAccountHandler(service)

	return &Module{
		Handler: handler,
		Service: service,
	}
}
