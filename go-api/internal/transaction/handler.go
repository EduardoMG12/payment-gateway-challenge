package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-gateway/go-api/internal/api"
	"payment-gateway/go-api/internal/i18n"
	"payment-gateway/go-api/internal/transaction/dto"

	"github.com/go-playground/validator/v10"
)

type TransactionHandler struct {
	service  TransactionService
	validate *validator.Validate
}

func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service:  service,
		validate: validator.New(),
	}
}

// @Summary Create a new transaction
// @Description Creates a new transaction in the payment gateway.
// @Accept json
// @Produce json
// @Param transaction body dto.CreateTransactionRequest true "Transaction data"
// @Success 201 {object} dto.ResponseCreateTransactionRequest
// @Failure 400 {object} api.APIError "Invalid request"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	var req dto.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Failed to decode request body:", err)
		api.WriteError(w, http.StatusBadRequest, i18n.GetErrorMessage(lang, i18n.ErrorInvalidRequestBody))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		api.WriteError(w, http.StatusBadRequest, i18n.GetErrorMessage(lang, i18n.ErrorInvalidRequestBody))
		return
	}

	createTx, err := h.service.CreateTransaction(r.Context(), req)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorCreatingTransaction))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createTx)
}

// @Summary Get transactions by Account ID
// @Description Returns a list of transactions for a specific account.
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {array} dto.ResponseCreateTransactionRequest
// @Failure 404 "Account not found"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /transactions/{accountId} [get]
func (h *TransactionHandler) GetTransactionByAccountId(w http.ResponseWriter, r *http.Request) {}
