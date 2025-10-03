package models

import "database/sql"

// NullableString represents a string value that may be null.
// Commonly used in database fields where null values are allowed.
type NullableString struct {
	// @Description Actual string value. Empty if Valid is false.
	// @Example "some-value"
	String string `json:"string"`

	// @Description Indicates whether the value is valid (true) or null (false).
	// @Example true
	Valid bool `json:"valid"`
}

// Transaction represents a financial transaction in the system.
type Transaction struct {
	// @Description Unique identifier for the transaction (UUID).
	// @Format uuid
	// @Example a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
	ID string `json:"id" db:"id"`

	// @Description Identifier of the account associated with this transaction (UUID).
	// @Format uuid
	// @Example e8b4d4c2-f9b6-4b1e-8e5e-9a9c2c1a1a9e
	AccountId string `json:"account_id" db:"account_id"`

	// @Description Identifier of the card used for the transaction. Nullable.
	// @Format uuid
	// @Example f0c3a2a6-0b3c-4a3e-8c7a-5b12bf7e4e1a
	CardId sql.NullString `json:"card_id" db:"card_id" swaggertype:"string" extensions:"x-nullable"`

	// @Description Identifier of the original transaction when this is a refund. Nullable.
	// @Format uuid
	// @Example c7a3c3b1-a2e4-4a25-8c7a-5b12bf7e4e1a
	RefundTransactionId sql.NullString `json:"refund_transaction_id" db:"refund_transaction_id" swaggertype:"string" extensions:"x-nullable"`

	// @Description Transaction amount in the smallest currency unit (e.g., cents). Must be positive.
	// @Minimum 1
	// @Example 5000
	AmountCents int64 `json:"amount_cents" db:"amount_cents"`

	// @Description Current status of the transaction.
	// @Enum PENDING APPROVED REJECTED
	// @Example PENDING
	Status string `json:"status" db:"status"`

	// @Description Type of the transaction.
	// @Enum DEPOSIT PURCHASE REFUND
	// @Example DEPOSIT
	Type string `json:"type" db:"type"`

	// @Description Unique key to guarantee idempotency of the transaction.
	// @Example 2025-10-03-17:30:00:e8b4d4c2:DEPOSIT:5000
	IdempotencyKey string `json:"idempotency_key" db:"idempotency_key"`

	// @Description Timestamp when the transaction was created (UTC, RFC3339 format).
	// @Format date-time
	// @Example 2025-10-03T20:30:00.123Z
	CreatedAt string `json:"created_at" db:"created_at"`
}
