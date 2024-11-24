package models

import "time"

type Ticker struct {
	Symbol    string    `json:"symbol"`
	LastPrice float64   `json:"last_price"`
	High24h   float64   `json:"high_24h"`
	Low24h    float64   `json:"low_24h"`
	Volume24h float64   `json:"volume_24h"`
	Change    float64   `json:"change"`
	Timestamp time.Time `json:"timestamp"`
}
