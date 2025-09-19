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
	GetTransactionByID(ctx context.Context, txID string) (*models.Transaction, error)
	CheckIdempotencyKey(ctx context.Context, key string) (bool, error)
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
		RETURNING id
	`
	var cardID sql.NullString
	if tx.CardID.Valid {
		cardID = tx.CardID
	}

	err := r.db.QueryRowContext(
		ctx,
		query,
		tx.AccountID,
		cardID,
		tx.AmountCents,
		"PENDING",
		tx.Type,
		tx.IdempotencyKey,
		time.Now().UTC(),
	).Scan(&tx.ID)

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

func (r *transactionRepositoryImpl) CheckIdempotencyKey(ctx context.Context, key string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM transactions WHERE idempotency_key = $1)`
	var exists bool

	err := r.db.QueryRowContext(ctx, query, key).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check idempotency key: %w", err)
	}

	return exists, nil
}
