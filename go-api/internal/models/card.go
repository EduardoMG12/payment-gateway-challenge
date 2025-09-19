package models

type Card struct {
	ID        string `json:"id" db:"id"`
	AccountID string `json:"account_id" db:"account_id"`
	CardToken string `json:"card_token" db:"card_token"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
