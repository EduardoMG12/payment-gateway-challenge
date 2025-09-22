package models

type Account struct {
	// @Description The unique identifier of the account.
	// @Example 550e8400-e29b-41d4-a716-446655440000
	ID string `json:"id" db:"id"`

	// @Description The unique username for the account.
	// @Example charlie
	Username string `json:"username" db:"username"`

	// @Description The creation timestamp of the account.
	// @Example 2025-09-22T19:15:24.526505Z
	CreatedAt string `json:"created_at" db:"created_at"`

	// @Description The last update timestamp of the account.
	// @Example 2025-09-22T19:15:24.526505Z
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
