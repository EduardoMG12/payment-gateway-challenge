package transaction

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"payment-gateway/go-api/internal/account"
	"payment-gateway/go-api/internal/card"
	"payment-gateway/go-api/internal/connection"

	"payment-gateway/go-api/internal/models"
	"payment-gateway/go-api/internal/repository"
	"payment-gateway/go-api/internal/transaction/dto"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, tx dto.CreateTransactionRequest) (*models.Transaction, error)
	GetBalanceByAccountId(ctx context.Context, accountId string) error
	GetBalanceFromCache(ctx context.Context, key string) (string, error)
	GetAllTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error)
	GetAllTransactionsByCardId(ctx context.Context, cardId string) ([]*models.Transaction, error)
	FindTransactionById(ctx context.Context, transactionId string) (*models.Transaction, error)
}

type transactionServiceImpl struct {
	repo           repository.TransactionRepository
	accountService account.AccountService
	cardService    card.CardService
	mqClient       connection.RabbitMQClient
	redis          connection.RedisConnection
}

func NewTransactionService(repo repository.TransactionRepository, service account.AccountService, mqClient connection.RabbitMQClient, cardService card.CardService, redis connection.RedisConnection) *transactionServiceImpl {
	return &transactionServiceImpl{repo: repo, accountService: service, mqClient: mqClient, cardService: cardService, redis: redis}
}

func (s *transactionServiceImpl) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (*models.Transaction, error) {
	timePrefix := time.Now().Format("2006-01-02-15:04:05.000")
	idempotencyKey := fmt.Sprintf(
		"%s:%s:%s:%d",
		timePrefix,
		req.AccountId,
		req.Type,
		req.AmountCents,
	)

	existingTx, err := s.repo.FindMostRecentTransaction(ctx, req.AccountId, req.Type, req.AmountCents)
	if err != nil {
		return nil, err
	}

	var parsedTime time.Time

	if existingTx != nil {
		parsedTime, _ = time.Parse(time.RFC3339, existingTx.CreatedAt)
	}
	if (time.Since(parsedTime) <= 3*time.Minute) && existingTx != nil {
		return existingTx, nil
	}

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin database transaction: %w", err)
	}
	defer tx.Rollback()

	account, err := s.accountService.GetAccountById(ctx, req.AccountId)
	if err != nil {
		return nil, err
	}
	var cardId string

	if req.CardToken != nil {
		cardIsReturn, err := s.cardService.GetCardByTokenAndAccountId(ctx, *req.CardToken, req.AccountId)
		if err != nil {
			return nil, err
		}
		cardId = cardIsReturn
	}

	if req.Type == "REFUND" {
		_, err = s.repo.GetTransactionByID(ctx, *req.RefundTransactionId)
		if err != nil {
			return nil, fmt.Errorf("original transaction for refund not found: %w", err)
		}
	}

	transaction := &models.Transaction{
		AccountId:           account.ID,
		CardId:              sql.NullString{String: "", Valid: false},
		RefundTransactionId: sql.NullString{String: "", Valid: false},
		AmountCents:         req.AmountCents,
		Type:                req.Type,
		IdempotencyKey:      idempotencyKey,
	}
	if cardId != "" {
		transaction.CardId = sql.NullString{String: cardId, Valid: true}
	}
	if req.RefundTransactionId != nil && req.Type == "REFUND" {
		transaction.RefundTransactionId = sql.NullString{String: *req.RefundTransactionId, Valid: true}
	}

	if err := s.repo.CreateTransaction(ctx, transaction); err != nil {
		return nil, fmt.Errorf("fail to create transaction: %w", err)
	}

	message, err := json.Marshal(transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize transaction for queue: %w", err)
	}

	if err := s.mqClient.Publish(ctx, "transactions_queue", message); err != nil {
		return nil, fmt.Errorf("failed to publish message to RabbitMQ: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit database transaction: %w", err)
	}

	return transaction, nil
}

func (s *transactionServiceImpl) GetBalanceFromCache(ctx context.Context, key string) (string, error) {
	return s.redis.Client.Get(ctx, key).Result()
}

func (s *transactionServiceImpl) GetBalanceByAccountId(ctx context.Context, accountId string) error {
	message := map[string]string{"account_id": accountId}

	account, err := s.accountService.GetAccountById(ctx, accountId)
	if err != nil {
		return fmt.Errorf("failed to get account by ID: %w", err)
	}
	if account == nil {
		return fmt.Errorf("account not found")
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize message for queue: %w", err)
	}

	return s.mqClient.Publish(ctx, "calculate_balance_queue", messageBytes)
}

func (s *transactionServiceImpl) GetAllTransactionsByAccountId(ctx context.Context, accountId string) ([]*models.Transaction, error) {
	err, transactions := s.repo.GetAllTransactionsByAccountIdTest(ctx, accountId)
	if err != nil {
		return nil, err
	}
	if len(transactions) > 0 {
		return transactions, nil
	}
	return transactions, nil
}

func (s *transactionServiceImpl) GetAllTransactionsByCardId(ctx context.Context, cardId string) ([]*models.Transaction, error) {
	transactions, err := s.repo.GetAllTransactionsByCardId(ctx, cardId)
	if err != nil {
		return nil, err
	}
	if len(transactions) > 0 {
		return transactions, nil
	}
	return []*models.Transaction{}, nil
}

func (s *transactionServiceImpl) FindTransactionById(ctx context.Context, transactionId string) (*models.Transaction, error) {
	transaction, err := s.repo.FindTransactionById(ctx, transactionId)
	if err != nil {
		return nil, err
	}
	if transaction != nil {
		return transaction, nil
	}
	return nil, nil
}
