package services

import (
	"pulsefin/database"
	"pulsefin/models"
)

func GetWatchList() ([]models.Watchlist, error) {
	var watchList []models.Watchlist

	query := "SELECT * FROM watchlist"
	err := database.DB.Select(&watchList, query)

	return watchList, err
}
