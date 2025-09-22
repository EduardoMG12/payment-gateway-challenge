package dto

type CreateAccountRequest struct {
	// @Description The username of the new account.
	// @Example charlie
	Username string `json:"username" validate:"required,min=3,max=100"`
}
