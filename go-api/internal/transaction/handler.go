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

func (h *TransactionHandler) GetBalanceByAccountId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	// Chave do Redis
	redisKey := "balance:" + accountId

	// 1. Tentar buscar no Redis
	balance, err := h.service.GetBalanceFromCache(ctx, redisKey)

	if err == redis.Nil { // redis.Nil significa "chave não encontrada" (Cache Miss)
		// 2. Chave não encontrada: Enviar para a fila
		err := h.service.GetBalanceByAccountId(ctx, accountId)
		if err != nil {
			api.WriteError(w, http.StatusInternalServerError, "Failed to request balance calculation")
			return
		}

		// 3. Responder com 202 Accepted
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "processing",
			"message": "The account balance is being calculated. Please try again in a few moments.",
		})
		return
	} else if err != nil { // Outro erro do Redis
		api.WriteError(w, http.StatusInternalServerError, "Error fetching balance from cache")
		return
	}

	// 4. Sucesso (Cache Hit): Retornar o saldo encontrado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"account_id":    accountId,
		"balance_cents": balance,
	})
}

func (h *TransactionHandler) GetAllTransactionByAccountIdTestOrderDate(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	vars := mux.Vars(r)
	accountId := vars["accountId"]

	err, transactions := h.service.GetAllTransactionsByAccountIdTest(r.Context(), accountId)
	fmt.Print("teste")
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorCreatingTransaction))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)

}
