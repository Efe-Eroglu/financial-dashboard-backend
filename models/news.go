package models

import "time"

type News struct {
	ID          int       `json:"id" db:"id"`
	StockSymbol string    `json:"stock_symbol" db:"stock_symbol"`
	Headline    string    `json:"headline" db:"headline"`
	Content     string    `json:"content" db:"content"`
	Source      string    `json:"source" db:"source"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
}
