package services

import (
	"errors"
	"pulsefin/database"
	"pulsefin/models"
)

func GetWatchList() ([]models.Watchlist, error) {
	var watchList []models.Watchlist

	query := "SELECT * FROM watchlist"
	err := database.DB.Select(&watchList, query)

	return watchList, err
}

func AddToWatchlist(userID int, stockSymbol string) error {
	// Hisse zaten izleme listesinde mi kontrol et
	var existingWatchlist models.Watchlist
	query := "SELECT id FROM watchlist WHERE user_id = $1 AND stock_symbol = $2"
	err := database.DB.Get(&existingWatchlist, query, userID, stockSymbol)
	if err == nil {
		return errors.New("stock already in watchlist")
	}

	// İzleme listesine hisse ekle
	query = "INSERT INTO watchlist (user_id, stock_symbol) VALUES ($1, $2)"
	_, err = database.DB.Exec(query, userID, stockSymbol)
	if err != nil {
		return errors.New("failed to add stock to watchlist")
	}

	return nil
}

func DeleteToWatchlist(userID int, stockSymbol string) error {
	var count int

	query := "SELECT COUNT(*) FROM watchlist WHERE user_id = $1 AND stock_symbol = $2"

	err := database.DB.Get(&count, query, userID, stockSymbol)
	if err != nil || count == 0 {
		return errors.New("stock not found in watchlist")
	}

	// Hisseyi izleme listesinden sil
	query = "DELETE FROM watchlist WHERE user_id = $1 AND stock_symbol = $2"
	_, err = database.DB.Exec(query, userID, stockSymbol)
	if err != nil {
		return errors.New("failed to remove stock from watchlist")
	}

	return nil
}
