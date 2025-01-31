package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"fin-api/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository реализует интерфейс repository.Repository для тестирования
type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) GetUser(ctx context.Context, userID int) (*model.User, error) {
    args := m.Called(ctx, userID)
    return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockRepository) UpdateUserBalance(ctx context.Context, userID int, amount float64) error {
    args := m.Called(ctx, userID, amount)
    return args.Error(0)
}

func (m *MockRepository) CreateTransaction(ctx context.Context, fromUserID, toUserID int, amount float64) error {
    args := m.Called(ctx, fromUserID, toUserID, amount)
    return args.Error(0)
}

func (m *MockRepository) GetLastTransactions(ctx context.Context, userID int, limit int) ([]*model.Transaction, error) {
    args := m.Called(ctx, userID, limit)
    return args.Get(0).([]*model.Transaction), args.Error(1)
}

func TestTopUpBalance(t *testing.T) {
    repo := new(MockRepository)
    srv := NewService(repo)

    tests := []struct {
        name     string
        userID   int
        amount   float64
        expected error
    }{
        {"InvalidAmount", 1, -100, errors.New("invalid amount")},
        {"ValidAmount", 1, 100, nil},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            repo.On("UpdateUserBalance", mock.Anything, test.userID, test.amount).Return(test.expected)
            err := srv.TopUpBalance(context.Background(), test.userID, test.amount)
            assert.Equal(t, test.expected, err)
            repo.AssertExpectations(t)
        })
    }
}

func TestTransferMoney(t *testing.T) {
    repo := new(MockRepository)
    srv := NewService(repo)

    tests := []struct {
        name         string
        fromUserID   int
        toUserID     int
        amount       float64
        expected     error
        userBalance  float64
    }{
        {"InvalidAmount", 1, 2, -100, errors.New("invalid amount"), 0},
        {"SameUser", 1, 1, 100, errors.New("cannot transfer money to yourself"), 0},
        {"InsufficientFunds", 1, 2, 100, errors.New("insufficient funds"), 50},
        {"ValidTransfer", 1, 2, 100, nil, 200},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            repo.On("GetUser", mock.Anything, test.fromUserID).Return(&model.User{ID: test.fromUserID, Balance: test.userBalance}, nil)
            repo.On("CreateTransaction", mock.Anything, test.fromUserID, test.toUserID, test.amount).Return(test.expected)
            err := srv.TransferMoney(context.Background(), test.fromUserID, test.toUserID, test.amount)
            assert.Equal(t, test.expected, err)
            repo.AssertExpectations(t)
        })
    }
}

func TestGetLastTransactions(t *testing.T) {
    repo := new(MockRepository)
    srv := NewService(repo)

    tests := []struct {
        name         string
        userID       int
        expected     []*model.Transaction
        expectedErr  error
    }{
        {"NoTransactions", 1, []*model.Transaction{}, nil},
        {"WithTransactions", 1, []*model.Transaction{
            {ID: 1, FromUserID: 1, ToUserID: 2, Amount: 100, CreatedAt: time.Now()},
        }, nil},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            repo.On("GetLastTransactions", mock.Anything, test.userID, 10).Return(test.expected, test.expectedErr)
            actual, err := srv.GetLastTransactions(context.Background(), test.userID)
            assert.Equal(t, test.expected, actual)
            assert.Equal(t, test.expectedErr, err)
            repo.AssertExpectations(t)
        })
    }
}