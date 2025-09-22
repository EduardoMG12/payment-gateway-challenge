package account

import (
	"encoding/json"
	"net/http"
	"payment-gateway/go-api/internal/account/dto"
	"payment-gateway/go-api/internal/api"
	"payment-gateway/go-api/internal/i18n"
	"payment-gateway/go-api/internal/models"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	service  AccountService
	validate *validator.Validate
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{
		service:  service,
		validate: validator.New(),
	}
}

// @Summary Create a new account
// @Description Creates a new account with a username.
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body dto.CreateAccountRequest true "Account data for creation"
// @Success 201 {object} models.Account
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	var req dto.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, http.StatusBadRequest, i18n.GetErrorMessage(lang, i18n.ErrorInvalidRequestBody))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.service.CreateAccount(r.Context(), req.Username)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorFailedToCreateAccount))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// @Summary Get all accounts with pagination
// @Description Returns a list of all accounts, with pagination support.
// @Tags accounts
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} models.Account
// @Router /accounts [get]
func (h *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	if limit > 17 {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.PaginationLimitExceeded))
		return
	}

	accounts, err := h.service.GetAllAccounts(r.Context(), page, limit)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorFailedToGetAccounts))
		return
	}
	if accounts == nil {
		accounts = make([]*models.Account, 0)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
