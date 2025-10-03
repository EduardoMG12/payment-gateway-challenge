package models

// Card represents a payment card linked to an account.
type Card struct {
	// @Description Unique identifier of the card (UUID).
	// @Format uuid
	// @Example 550e8400-e29b-41d4-a716-446655440000
	ID string `json:"id" db:"id"`

	// @Description Unique identifier of the account this card belongs to (UUID).
	// @Format uuid
	// @Example 550e8400-e29b-41d4-a716-446655440000
	AccountId string `json:"account_id" db:"account_id"`

	// @Description Irreversible token representing the card. Used internally instead of raw card numbers.
	// @Format hash
	// @Example 5b7c16af7278094cd14bd041079111ed00fa832c8a460d8e3f40156408d99475
	CardToken string `json:"card_token" db:"card_token"`

	// @Description Last four digits of the card number, useful for display/identification.
	// @MinLength 4
	// @MaxLength 4
	// @Example 8995
	LastFourDigits string `json:"last_four_digits" db:"last_four_digits"`

	// @Description Timestamp when the card was created (UTC, RFC3339 format).
	// @Format date-time
	// @Example 2025-09-22T19:15:24.526505Z
	CreatedAt string `json:"created_at" db:"created_at"`
}
