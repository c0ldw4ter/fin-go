package repository

import (
	"context"
	"fmt"

	"fin-api/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository определяет интерфейс для работы с базой данных
type Repository interface {
    GetUser(ctx context.Context, userID int) (*model.User, error)
    UpdateUserBalance(ctx context.Context, userID int, amount float64) error
    CreateTransaction(ctx context.Context, fromUserID, toUserID int, amount float64) error
    GetLastTransactions(ctx context.Context, userID int, limit int) ([]*model.Transaction, error)
}

// Реализация интерфейса Repository
type repository struct {
    db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
    return &repository{db: db}
}

func (r *repository) GetUser(ctx context.Context, userID int) (*model.User, error) {
    var user model.User
    err := r.db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Balance)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return &user, nil
}

func (r *repository) UpdateUserBalance(ctx context.Context, userID int, amount float64) error {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    _, err = tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID)
    if err != nil {
        return fmt.Errorf("failed to update user balance: %w", err)
    }

    if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }
    return nil
}

func (r *repository) CreateTransaction(ctx context.Context, fromUserID, toUserID int, amount float64) error {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    _, err = tx.Exec(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromUserID)
    if err != nil {
        return fmt.Errorf("failed to update sender balance: %w", err)
    }

    _, err = tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toUserID)
    if err != nil {
        return fmt.Errorf("failed to update receiver balance: %w", err)
    }

    _, err = tx.Exec(ctx, "INSERT INTO transactions (from_user_id, to_user_id, amount) VALUES ($1, $2, $3)", fromUserID, toUserID, amount)
    if err != nil {
        return fmt.Errorf("failed to create transaction: %w", err)
    }

    if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }
    return nil
}

func (r *repository) GetLastTransactions(ctx context.Context, userID int, limit int) ([]*model.Transaction, error) {
    rows, err := r.db.Query(ctx, "SELECT id, from_user_id, to_user_id, amount, created_at FROM transactions WHERE from_user_id = $1 OR to_user_id = $1 ORDER BY created_at DESC LIMIT $2", userID, limit)
    if err != nil {
        return nil, fmt.Errorf("failed to get transactions: %w", err)
    }
    defer rows.Close()

    var transactions []*model.Transaction
    for rows.Next() {
        var transaction model.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.FromUserID, &transaction.ToUserID, &transaction.Amount, &transaction.CreatedAt); err != nil {
            return nil, fmt.Errorf("failed to scan transaction: %w", err)
        }
        transactions = append(transactions, &transaction)
    }

    return transactions, nil
}