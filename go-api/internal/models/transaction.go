package models

import "database/sql"

type Transaction struct {
	ID             string         `json:"id" db:"id"`
	AccountId      string         `json:"account_id" db:"account_id"`
	CardId         sql.NullString `json:"card_id" db:"card_id"`
	AmountCents    int64          `json:"amount_cents" db:"amount_cents"`
	Status         string         `json:"status" db:"status"`
	Type           string         `json:"type" db:"type"`
	IdempotencyKey string         `json:"idempotency_key" db:"idempotency_key"`
	CreatedAt      string         `json:"created_at" db:"created_at"`
}
