package service

import (
	"context"
	"errors"
	"fmt"

	"fin-api/internal/model"
	"fin-api/internal/repository"
)

type Service struct {
    repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) TopUpBalance(ctx context.Context, userID int, amount float64) error {
    if amount <= 0 {
        return errors.New("invalid amount")
    }
    return s.repo.UpdateUserBalance(ctx, userID, amount)
}

func (s *Service) TransferMoney(ctx context.Context, fromUserID, toUserID int, amount float64) error {
    if amount <= 0 {
        return errors.New("invalid amount")
    }
    if fromUserID == toUserID {
        return errors.New("cannot transfer money to yourself")
    }

    user, err := s.repo.GetUser(ctx, fromUserID)
    if err != nil {
        return fmt.Errorf("failed to get user: %w", err)
    }
    if user.Balance < amount {
        return errors.New("insufficient funds")
    }

    return s.repo.CreateTransaction(ctx, fromUserID, toUserID, amount)
}

func (s *Service) GetLastTransactions(ctx context.Context, userID int) ([]*model.Transaction, error) {
    return s.repo.GetLastTransactions(ctx, userID, 10)
}