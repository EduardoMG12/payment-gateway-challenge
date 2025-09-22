package dto

type CreateAccountRequest struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
}
