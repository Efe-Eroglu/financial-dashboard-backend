package models

import "time"

type Stock struct {
	ID                 int       `json:"id" db:"id"`
	Symbol             string    `json:"symbol" db:"symbol"`
	Name               string    `json:"name" db:"name"`
	PreviousClosePrice float64   `json:"previous_close_price" db:"previous_close_price"`
	HighPrice          float64   `json:"high_price" db:"high_price"`
	LowPrice           float64   `json:"low_price" db:"low_price"`
	Volume             int       `json:"volume" db:"volume"`
	MarketCap          float64   `json:"market_cap" db:"market_cap"`
	ChangePercentage   float64   `json:"change_percentage" db:"change_percentage"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}
