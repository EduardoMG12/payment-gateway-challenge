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
		INSER INTO cards (account_id, card_token) VALUES ($1, $2) RETURNING id;	
	`
	err := r.db.QueryRowContext(ctx, query, card.AccountID, card.CardToken).Scan(&card.ID)

	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}
