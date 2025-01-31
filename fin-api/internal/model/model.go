package model

import "time"

type User struct {
    ID      int     `json:"id"`
    Balance float64 `json:"balance"`
}

type Transaction struct {
    ID        int       `json:"id"`
    FromUserID int       `json:"from_user_id"`
    ToUserID   int       `json:"to_user_id"`
    Amount     float64   `json:"amount"`
    CreatedAt  time.Time `json:"created_at"`
}