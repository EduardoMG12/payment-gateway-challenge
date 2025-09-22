package repository

import (
	"context"
	"fmt"
	"payment-gateway/go-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type CardRepository interface {
	CreateCard(ctx context.Context, card *models.Card) error
}

type cardRepositoryImpl struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) CardRepository {
	return &cardRepositoryImpl{db: db}
}

func (r *cardRepositoryImpl) CreateCard(ctx context.Context, card *models.Card) error {
	query := `
        INSERT INTO cards (account_id, card_token, last_four_digits)
        VALUES ($1, $2, $3)
        RETURNING id, card_token, last_four_digits, created_at;
    `
	err := r.db.QueryRowContext(ctx, query, card.AccountID, card.CardToken, card.LastFourDigits).Scan(
		&card.ID,
		&card.CardToken,
		&card.LastFourDigits,
		&card.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create card: %w", err)
	}

	return nil
}
