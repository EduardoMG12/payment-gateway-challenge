package dto

type CreateTransactionRequest struct {
	AccountId   string  `json:"account_id" validate:"required,uuid4"`
	CardToken   *string `json:"card_token,omitempty" validate:"omitempty,min=20,max=126"`
	AmountCents int64   `json:"amount_cents" validate:"required,gt=0"`
	Type        string  `json:"type" validate:"required,oneof=DEPOSIT PURCHASE REFUND CHARGE"`
}

type ResponseCreateTransactionRequest struct {
	AccountId   string  `json:"account_id" validate:"required,uuid4"`
	CardId      *string `json:"card_id,omitempty" validate:"omitempty,uuid4"`
	AmountCents int64   `json:"amount_cents" validate:"required,gt=0"`
	Type        string  `json:"type" validate:"required,oneof=DEPOSIT PURCHASE REFUND CHARGE"`
}
