package repository

import (
	"context"
	"database/sql"
	"fmt"
	"payment-gateway/go-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *models.Account) error
	GetAllAccounts(ctx context.Context, page, limit int) ([]*models.Account, error)
	GetAccountById(ctx context.Context, id string) (*models.Account, error)
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
        RETURNING id, created_at, updated_at;
    `

	err := r.db.QueryRowContext(ctx, query, account.Username).Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}

func (r *accountRepositoryImpl) GetAllAccounts(ctx context.Context, page, limit int) ([]*models.Account, error) {
	offset := (page - 1) * limit

	query := `SELECT id, username, created_at, updated_at FROM accounts
        ORDER BY created_at DESC
		LIMIT $1 OFFSET $2; `

	var accounts []*models.Account
	err := r.db.SelectContext(ctx, &accounts, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}
	return accounts, nil
}

func (r *accountRepositoryImpl) GetAccountById(ctx context.Context, accountID string) (*models.Account, error) {
	query := `SELECT * FROM accounts WHERE id = $1`
	var account models.Account
	err := r.db.GetContext(ctx, &account, query, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return &account, nil
}
