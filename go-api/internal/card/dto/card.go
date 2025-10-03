package dto

// @Description Request body for creating a new card
type CreateCardRequest struct {
	// @Description Account ID to associate the new card (UUID)
	// @Example e252f5dd-ded2-4a30-a4a5-6e2940008d54
	AccountId string `json:"account_id" validate:"required,uuid4"`
}

// @Description Response when a card is created
type CardResponse struct {
	ID             string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AccountId      string `json:"account_id" example:"e252f5dd-ded2-4a30-a4a5-6e2940008d54"`
	CardToken      string `json:"card_token" example:"5b7c16af7278094cd14bd041079111ed00fa832c8a460d8e3f40156408d99475"`
	LastFourDigits string `json:"last_four_digits" example:"8995"`
	CreatedAt      string `json:"created_at" example:"2025-09-22T19:15:24.526505Z"`
}
