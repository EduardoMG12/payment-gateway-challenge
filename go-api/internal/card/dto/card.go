package dto

type CreateCardRequest struct {
	// @Description Account ID to associate the new card
	// @Example e252f5dd-ded2-4a30-a4a5-6e2940008d54
	AccountId string `json:"account_id" validate:"required"`
}
