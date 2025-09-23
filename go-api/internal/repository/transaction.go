package repository

import (
	"context"
	"database/sql"
	"fmt"
	"payment-gateway/go-api/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *models.Transaction) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
	GetTransactionByID(ctx context.Context, txID string) (*models.Transaction, error)
	FindMostRecentTransaction(ctx context.Context, accountID, txType string, amountCents int64) (*models.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) CreateTransaction(ctx context.Context, tx *models.Transaction) error {
	query := `
		INSERT INTO transactions (account_id, card_id, amount_cents, status, type, idempotency_key, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, status, created_at;
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		tx.AccountId,
		tx.CardId,
		tx.AmountCents,
		"PENDING",
		tx.Type,
		tx.IdempotencyKey,
		time.Now().UTC(),
	).Scan(&tx.ID, &tx.Status, &tx.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return nil
}

func (r *transactionRepositoryImpl) GetTransactionByID(ctx context.Context, txID string) (*models.Transaction, error) {
	query := `SELECT * FROM transactions WHERE id = $1`
	var tx models.Transaction

	err := r.db.GetContext(ctx, &tx, query, txID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get transaction by id: %w", err)
	}

	return &tx, nil
}

func (r *transactionRepositoryImpl) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("falha ao iniciar a transação: %w", err)
	}
	return tx, nil
}

func (r *transactionRepositoryImpl) FindMostRecentTransaction(ctx context.Context, accountID, txType string, amountCents int64) (*models.Transaction, error) {
	query := `
        SELECT *
        FROM transactions
        WHERE account_id = $1
        AND type = $2
        AND amount_cents = $3
        ORDER BY created_at DESC
        LIMIT 1
    `
	var tx models.Transaction

	err := r.db.GetContext(ctx, &tx, query, accountID, txType, amountCents)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find most recent transaction: %w", err)
	}

	return &tx, nil
}
