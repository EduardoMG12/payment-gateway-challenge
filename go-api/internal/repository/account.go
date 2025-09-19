package repository

import (
	"context"
	"fmt"
	"payment-gateway/go-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *models.Account) error
}

type accountRepositoryImpl struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepositoryImpl{db: db}
}

func (r *accountRepositoryImpl) CreateAccount(ctx context.Context, account *models.Account) error {
	query := `
        INSERT INTO accounts (username)
        VALUES ($1)
        RETURNING id;
    `

	err := r.db.QueryRowContext(ctx, query, account.Username).Scan(&account.ID)

	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}
