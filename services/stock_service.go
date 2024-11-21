package services

import (
	"errors"
	"pulsefin/database"
	"pulsefin/models"
)

func GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock

	query := "SELECT * FROM stocks"
	err := database.DB.Select(&stocks, query)

	return stocks, err
}

func GetStock(stockSymbol string) (*models.Stock, error) {
	var stock models.Stock

	query := "SELECT * FROM stocks WHERE symbol = $1"
	err := database.DB.Get(&stock, query, stockSymbol)

	if err != nil {
		return nil, errors.New("stock not foun")
	}

	return &stock, nil
}
