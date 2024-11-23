package models

import "time"

type Ticker struct {
	Symbol    string    `json:"symbol"`
	LastPrice float64   `json:"last_price"`
	High24h   float64   `json:"high_24h"`
	Low24h    float64   `json:"low_24h"`
	Volume24h float64   `json:"volume_24h"`
	Timestamp time.Time `json:"timestamp"`
}
