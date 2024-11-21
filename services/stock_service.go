package services

import (
	"pulsefin/database"
	"pulsefin/models"
)

func GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock

	query := "SELECT * FROM stocks"
	err := database.DB.Select(&stocks, query)

	return stocks, err
}
