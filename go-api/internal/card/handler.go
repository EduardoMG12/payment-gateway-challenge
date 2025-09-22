package card

import (
	"encoding/json"
	"net/http"
	"payment-gateway/go-api/internal/api"
	"payment-gateway/go-api/internal/card/dto"
	"payment-gateway/go-api/internal/i18n"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CardHandler struct {
	service  CardService
	validate *validator.Validate
}

func NewCardHandler(service CardService) *CardHandler {
	return &CardHandler{
		service:  service,
		validate: validator.New(),
	}
}

// @Summary Create a new card
// @Description Creates a new, fictional card and associates it with an account.
// @Tags cards
// @Accept json
// @Produce json
// @Param card body dto.CreateCardRequest true "Account ID to associate the new card"
// @Success 201 {object} models.Card
// @Failure 400 {object} api.APIError "Invalid request body or validation failed"
// @Failure 500 {object} api.APIError "Failed to create card"
// @Router /cards [post]
func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	var req dto.CreateCardRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, http.StatusBadRequest, i18n.GetErrorMessage(lang, i18n.ErrorInvalidRequestBody))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	card, err := h.service.CreateCard(r.Context(), req.AccountId)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorFailedToCreateCard))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

// @Summary Get all cards by account ID
// @Description Returns a list of all cards associated with an account, ordered by creation date.
// @Tags cards
// @Produce json
// @Param accountId path string true "ID of the account"
// @Success 200 {array} models.Card
// @Failure 400 {object} api.APIError "Invalid request or validation failed"
// @Failure 404 {object} api.APIError "Account not found"
// @Failure 500 {object} api.APIError "Failed to retrieve cards"
// @Router /cards/{accountId} [get]
func (h *CardHandler) GetAllCardsByAccountId(w http.ResponseWriter, r *http.Request) {
	lang := i18n.GetLangFromHeader(r)

	vars := mux.Vars(r)
	accountId := vars["accountId"]

	cards, err := h.service.GetAllCardsByAccountId(r.Context(), accountId)
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, i18n.GetErrorMessage(lang, i18n.ErrorToFindCards))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cards)
}
