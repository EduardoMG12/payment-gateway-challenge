package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-gateway/go-api/internal/api"
	"payment-gateway/go-api/internal/i18n"
	"payment-gateway/go-api/internal/transaction/dto"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
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

// @ID create-transaction
// @Summary Create a new transaction
// @Description Creates a new transaction (DEPOSIT, PURCHASE, REFUND, CHARGE) in the payment gateway.
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body dto.CreateTransactionRequest true "Transaction data"
// @Success 201 {object} dto.ResponseCreateTransactionRequest "Transaction created successfully"
// @Header 201 {string} Location "URL of the created transaction"
// @Failure 400 {object} api.APIError "Invalid request body or validation failed"
// @Failure 404 {object} api.APIError "Account not found"
// @Failure 422 {object} api.APIError "Business rule violation (e.g. insufficient funds)"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /transactions [post]
// @Example request {"account_id":"e7b40123-cb12-41fa-b5bc-5a128448027e","amount_cents":10000,"type":"PURCHASE","card_token":"16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6"}
// @Example response {"account_id":"e7b40123-cb12-41fa-b5bc-5a128448027e","card_id":"16ecac04-9e45-4a5b-b7d4-d6c1c66bafd6","amount_cents":10000,"type":"PURCHASE"}
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

// @ID get-transactions-by-account
// @Summary Get transactions by Account ID
// @Description Returns a list of transactions for a specific account.
// @Tags transactions
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {array} dto.ResponseCreateTransactionRequest
// @Failure 404 {object} api.APIError "Account not found"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /transactions/{accountId} [get]
func (h *TransactionHandler) GetTransactionByAccountId(w http.ResponseWriter, r *http.Request) {}

// @ID get-account-balance
// @Summary Get Account Balance
// @Description Retrieves the current balance for a specific account.
// If cached → returns immediately (200).
// If not cached → triggers background calc and returns processing (202).
// @Tags accounts
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} dto.ResponseAccountBalance "Balance retrieved from cache"
// @Success 202 {object} dto.ProcessingResponse "Balance calculation triggered"
// @Failure 404 {object} api.APIError "Account not found"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /accounts/{accountId}/balance [get]
func (h *TransactionHandler) GetBalanceByAccountId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	redisKey := "balance:" + accountId

	balance, err := h.service.GetBalanceFromCache(ctx, redisKey)

	if err == redis.Nil {

		err := h.service.GetBalanceByAccountId(ctx, accountId)
		if err != nil {
			api.WriteError(w, http.StatusInternalServerError, "Failed to request balance calculation")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "processing",
			"message": "The account balance is being calculated. Please try again in a few moments.",
		})
		return
	} else if err != nil {
		api.WriteError(w, http.StatusInternalServerError, "Error fetching balance from cache")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"account_id":    accountId,
		"balance_cents": balance,
	})
}

// @ID get-transactions-test
// @Summary Get All Transactions for an Account (Test)
// @Description Retrieves a list of all transactions for an account, ordered by creation date (desc).
// @Tags transactions
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {array} dto.ResponseCreateTransactionRequest
// @Failure 404 {object} api.APIError "Account not found"
// @Failure 500 {object} api.APIError "Internal server error"
// @Router /transactions/test/{accountId} [get]
func (h *TransactionHandler) GetAllTransactionByAccountIdTestOrderDate(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	vars := mux.Vars(r)
	accountId := vars["accountId"]

	transactions, err := h.service.GetAllTransactionsByAccountIdTest(r.Context(), accountId)
	fmt.Print("teste")
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorFindAllTransaction))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)

}
