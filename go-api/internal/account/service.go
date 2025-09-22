package account

import (
	"context"
	"payment-gateway/go-api/internal/models"
	"payment-gateway/go-api/internal/repository"
)

type AccountService interface {
	CreateAccount(ctx context.Context, username string) (*models.Account, error)
	GetAllAccounts(ctx context.Context, page, limit int) ([]*models.Account, error)
}

type accountServiceImpl struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *accountServiceImpl {
	return &accountServiceImpl{repo: repo}
}

func (s *accountServiceImpl) CreateAccount(ctx context.Context, username string) (*models.Account, error) {
	account := &models.Account{
		Username: username,
	}
	if err := s.repo.CreateAccount(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *accountServiceImpl) GetAllAccounts(ctx context.Context, page, limit int) ([]*models.Account, error) {
	accounts, err := s.repo.GetAllAccounts(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
