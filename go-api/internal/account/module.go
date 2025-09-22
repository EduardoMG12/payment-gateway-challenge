package account

import (
	"payment-gateway/go-api/internal/repository"

	"github.com/jmoiron/sqlx"
)

// Module representa o módulo de conta completo.
type Module struct {
	Handler *AccountHandler
	Service AccountService
}

// NewModule cria e retorna um novo módulo de conta, injetando suas dependências.
func NewModule(db *sqlx.DB) *Module {
	repo := repository.NewAccountRepository(db)
	service := NewAccountService(repo)
	handler := NewAccountHandler(service)

	return &Module{
		Handler: handler,
		Service: service,
	}
}
