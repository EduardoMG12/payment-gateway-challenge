package dto

type CreateTransactionRequest struct {
	// @Description The account's ID for which the transaction will be performed.
	// @Example e7b40123-cb12-41fa-b5bc-5a128448027e
	AccountId string `json:"account_id" validate:"required,uuid4"`

	// @Description The credit card token (optional for some transaction types like DEPOSIT).
	// @Example 16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6
	CardToken *string `json:"card_token,omitempty" validate:"omitempty,min=20,max=126"`

	// @Description The ID of the transaction being refunded (required for REFUND transactions).
	// @Example 16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6
	RefundTransactionId *string `json:"refund_transaction_id,omitempty" validate:"omitempty,uuid4"`

	// @Description The transaction amount in cents. Must be a positive integer.
	// @Example 10000
	AmountCents int64 `json:"amount_cents" validate:"required,gt=0"`

	// @Description The type of transaction. Valid options are: DEPOSIT, PURCHASE, REFUND, CHARGE.
	// @Example PURCHASE
	Type string `json:"type" validate:"required,oneof=DEPOSIT PURCHASE REFUND CHARGE"`
}

type ResponseCreateTransactionRequest struct {
	// @Description The ID of the account associated with the transaction.
	// @Example e7b40123-cb12-41fa-b5bc-5a128448027e
	AccountId string `json:"account_id" validate:"required,uuid4"`

	// @Description The ID of the credit card used for the transaction. It will be null if no card was associated with the transaction.
	// @Example 16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6
	CardId *string `json:"card_id,omitempty" validate:"omitempty,uuid4"`

	// @Description The transaction amount in cents.
	// @Example 10000
	AmountCents int64 `json:"amount_cents" validate:"required,gt=0"`

	// @Description The type of the transaction.
	// @Example PURCHASE
	Type string `json:"type" validate:"required,oneof=DEPOSIT PURCHASE REFUND CHARGE"`
}

type ResponseAccountBalance struct {
	// @Description The unique identifier of the account.
	// @Example 550e8400-e29b-41d4-a716-446655440000
	Id string `json:"id" db:"id"`

	// @Description The all sum of transactions amount in cents. Must be a positive integer.
	// @Example 10000
	Balance int64 `json:"balance_cents"`
}
type RequestGetAccountAndBalance struct {
	// @Description The unique identifier of the account.
	// @Example 550e8400-e29b-41d4-a716-446655440000
	Id string `json:"id" db:"id"`
}
