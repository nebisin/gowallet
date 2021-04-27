package model

import "time"

type Account struct {
	ID        uint64    `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        uint64    `json:"id"`
	AccountID uint64    `json:"account_id"`
	Amount    int64     `json:"amount"` // can be negative or positive
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            uint64    `json:"id"`
	FromAccountID uint64    `json:"from_account_id"`
	ToAccountID   uint64    `json:"to_account_id"`
	Amount        int64     `json:"amount"` // must be positive
	CreatedAt     time.Time `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
}
