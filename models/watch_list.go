package models

import "time"

type Watchlist struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	StockSymbol string    `json:"stock_symbol" db:"stock_symbol"`
	AddedAt     time.Time `json:"added_at" db:"added_at"`
}
