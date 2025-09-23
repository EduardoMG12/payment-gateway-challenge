package card

import (
	"context"
	"fmt"
	"payment-gateway/go-api/internal/models"
	"payment-gateway/go-api/internal/repository"
	"payment-gateway/go-api/internal/utils"

	"payment-gateway/go-api/internal/account"
)

type CardService interface {
	CreateCard(ctx context.Context, accountID string) (*models.Card, error)
	GetAllCardsByAccountId(ctx context.Context, accountId string) ([]*models.Card, error)
	GetCardByTokenAndAccountId(ctx context.Context, cardToken, accountId string) (string, error)
}

type cardServiceImpl struct {
	repo           repository.CardRepository
	accountService account.AccountService
}

func NewCardService(repo repository.CardRepository, accountService account.AccountService) *cardServiceImpl {
	return &cardServiceImpl{repo: repo, accountService: accountService}
}

func (s *cardServiceImpl) CreateCard(ctx context.Context, accountID string) (*models.Card, error) {

	account, err := s.accountService.GetAccountById(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	cardNumber, cvc, expiryMonth, expiryYear, err := utils.GenerateCardDetails()
	if err != nil {
		return nil, fmt.Errorf("failed to generate card details: %w", err)
	}

	cardToken := utils.GenerateCardToken(cardNumber, cvc, expiryMonth, expiryYear)

	lastFourDigits := cardNumber[len(cardNumber)-4:]

	card := &models.Card{
		AccountId:      account.ID,
		CardToken:      cardToken,
		LastFourDigits: lastFourDigits,
	}

	if err := s.repo.CreateCard(ctx, card); err != nil {
		return nil, fmt.Errorf("failed to create card: %w", err)
	}

	return card, nil
}

func (s *cardServiceImpl) GetAllCardsByAccountId(ctx context.Context, accountId string) ([]*models.Card, error) {
	account, err := s.accountService.GetAccountById(ctx, accountId)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account not found")
	}

	cards, err := s.repo.GetAllCardsByAccountId(ctx, account.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}

	return cards, nil
}

func (s *cardServiceImpl) GetCardByTokenAndAccountId(ctx context.Context, cardToken, accountId string) (string, error) {
	card, err := s.repo.GetCardByTokenAndAccountId(ctx, cardToken, accountId)

	if err != nil {
		return "", fmt.Errorf("failed to get card by token and account id: %w", err)
	}

	if card == "" {
		return "", fmt.Errorf("card not found")
	}

	return card, nil
}
