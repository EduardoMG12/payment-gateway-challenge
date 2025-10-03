package models

// Account represents a user account in the system.
type Account struct {
	// @Description Unique identifier of the account (UUID).
	// @Format uuid
	// @Example 550e8400-e29b-41d4-a716-446655440000
	ID string `json:"id" db:"id"`

	// @Description Unique username chosen by the account owner. Must be unique across the system.
	// @MinLength 3
	// @MaxLength 32
	// @Example charlie
	Username string `json:"username" db:"username"`

	// @Description Timestamp when the account was created (UTC, RFC3339 format).
	// @Format date-time
	// @Example 2025-09-22T19:15:24.526505Z
	CreatedAt string `json:"created_at" db:"created_at"`

	// @Description Timestamp of the last account update (UTC, RFC3339 format).
	// @Format date-time
	// @Example 2025-09-22T19:15:24.526505Z
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
