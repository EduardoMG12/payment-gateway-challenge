package dto

// @Description Request body for creating a new transaction
type CreateTransactionRequest struct {
	// @Description The account's ID for which the transaction will be performed (UUID).
	AccountId string `json:"account_id" validate:"required,uuid4" example:"e7b40123-cb12-41fa-b5bc-5a128448027e"`

	// @Description The credit card token (optional for some transaction types like DEPOSIT).
	CardToken *string `json:"card_token,omitempty" validate:"omitempty,min=20,max=126" example:"16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6"`

	// @Description The ID of the transaction being refunded (only for REFUND).
	RefundTransactionId *string `json:"refund_transaction_id,omitempty" validate:"omitempty,uuid4" example:"3c2b4791-7f84-4d77-b2e0-56de8df97f33"`

	// @Description Transaction amount in cents. Must be positive.
	AmountCents int64 `json:"amount_cents" validate:"required,gt=0" example:"10000"`

	// @Description Transaction type: DEPOSIT, PURCHASE, REFUND, CHARGE
	Type string `json:"type" validate:"required,oneof=DEPOSIT PURCHASE REFUND CHARGE" example:"PURCHASE"`
}

// @Description Response returned when a transaction is created or queried
type ResponseCreateTransactionRequest struct {
	AccountId   string  `json:"account_id" example:"e7b40123-cb12-41fa-b5bc-5a128448027e"`
	CardId      *string `json:"card_id,omitempty" example:"16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6"`
	AmountCents int64   `json:"amount_cents" example:"10000"`
	Type        string  `json:"type" example:"PURCHASE"`
}

// @Description Response for account balance
type ResponseAccountBalance struct {
	Id      string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Balance int64  `json:"balance_cents" example:"10000"`
}

// @Description Response when balance calculation is processing
type ProcessingResponse struct {
	Status  string `json:"status" example:"processing"`
	Message string `json:"message" example:"The account balance is being calculated. Please try again later."`
}
