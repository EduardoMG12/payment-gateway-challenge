package models

type Card struct {
	// @Description The unique identifier of the card.
	// @Example 550e8400-e29b-41d4-a716-446655440000
	ID string `json:"id" db:"id"`

	// @Description The unique identifier of the account this card belongs to.
	// @Example 550e8400-e29b-41d4-a716-446655440000
	AccountId string `json:"account_id" db:"account_id"`

	// @Description The unique, irreversible token for the card.
	// @Example 5b7c16af7278094cd14bd041079111ed00fa832c8a460d8e3f40156408d99475
	CardToken string `json:"card_token" db:"card_token"`

	// @Description The last four digits of the card number for identification.
	// @Example 8995
	LastFourDigits string `json:"last_four_digits" db:"last_four_digits"`

	// @Description The creation timestamp of the card.
	// @Example 2025-09-22T19:15:24.526505Z
	CreatedAt string `json:"created_at" db:"created_at"`
}
