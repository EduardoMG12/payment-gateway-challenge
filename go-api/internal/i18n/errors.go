package i18n

import (
	"net/http"
	"strings"
)

const (
	ErrorInvalidRequestBody        = "invalid_request_body"
	ErrorMethodNotAllowed          = "method_not_allowed"
	ErrorInternalServerError       = "internal_server_error"
	ErrorValidationFailed          = "validation_failed"
	ErrorNotFound                  = "not_found"
	ErrorAccountNotFound           = "account_not_found"
	ErrorCardNotFound              = "card_not_found"
	ErrorTransactionNotFound       = "transaction_not_found"
	ErrorDuplicateKey              = "duplicate_key"
	ErrorInsufficientFunds         = "insufficient_funds"
	ErrorFailedToCreateAccount     = "failed_to_create_account"
	ErrorFailedToCreateCard        = "failed_to_create_card"
	ErrorFailedToCreateTransaction = "failed_to_create_transaction"
	ErrorFailedToGetAccounts       = "failed_to_get_accounts"
	PaginationLimitExceeded        = "pagination_limit_exceeded"
	ErrorToFindCards               = "error_to_find_cards"
)

var errorMessages = map[string]map[string]string{
	"en-us": {
		ErrorInvalidRequestBody:        "Invalid request body",
		ErrorMethodNotAllowed:          "Method not allowed",
		ErrorInternalServerError:       "Internal server error",
		ErrorValidationFailed:          "Validation failed",
		ErrorNotFound:                  "Resource not found",
		ErrorAccountNotFound:           "Account not found",
		ErrorCardNotFound:              "Card not found",
		ErrorTransactionNotFound:       "Transaction not found",
		ErrorDuplicateKey:              "Duplicate key",
		ErrorInsufficientFunds:         "Insufficient funds",
		ErrorFailedToCreateAccount:     "Failed to create account",
		ErrorFailedToCreateCard:        "Failed to create card",
		ErrorFailedToCreateTransaction: "Failed to create transaction",
		ErrorFailedToGetAccounts:       "Failed to get accounts",
		PaginationLimitExceeded:        "Pagination limit exceeded",
		ErrorToFindCards:               "Error to find cards",
	},
	"pt-br": {
		ErrorInvalidRequestBody:        "Corpo da requisição inválido",
		ErrorMethodNotAllowed:          "Método não permitido",
		ErrorInternalServerError:       "Erro interno do servidor",
		ErrorValidationFailed:          "Falha na validação",
		ErrorNotFound:                  "Recurso não encontrado",
		ErrorAccountNotFound:           "Conta não encontrada",
		ErrorCardNotFound:              "Cartão não encontrado",
		ErrorTransactionNotFound:       "Transação não encontrada",
		ErrorDuplicateKey:              "Chave duplicada",
		ErrorInsufficientFunds:         "Saldo insuficiente",
		ErrorFailedToCreateAccount:     "Falha ao criar a conta",
		ErrorFailedToCreateCard:        "Falha ao criar o cartão",
		ErrorFailedToCreateTransaction: "Falha ao criar a transação",
		ErrorFailedToGetAccounts:       "Falha ao buscar contas",
		PaginationLimitExceeded:        "Limite de paginação excedido",
		ErrorToFindCards:               "Erro ao buscar cartões",
	},
}

func GetErrorMessage(lang, key string) string {
	lang = strings.ToLower(lang)
	if messages, ok := errorMessages[lang]; ok {
		if msg, ok := messages[key]; ok {
			return msg
		}
	}

	return errorMessages["en-US"][key]
}

func GetLangFromHeader(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		return "en-US"
	}

	parts := strings.Split(lang, ",")
	return strings.TrimSpace(parts[0])
}
