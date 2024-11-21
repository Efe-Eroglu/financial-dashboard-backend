package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

func GetWatchList(c echo.Context) error {

	userID := c.Get("userID").(int)

	watchlist, err := services.GetWatchlist(userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch watclist"})
	}

	return c.JSON(http.StatusOK, watchlist)

}

func AddToWatchlist(c echo.Context) error {
	userID := c.Get("userID").(int)

	var request struct {
		StockSymbol string `json:"stock_symbol"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// İzleme listesine hisse ekleme işlemi
	err := services.AddToWatchlist(userID, request.StockSymbol)
	if err != nil {
		if err.Error() == "stock already in watchlist" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Stock already in watchlist"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add stock to watchlist"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Stock added to watchlist successfully"})
}

func DeleteToWatchlist(c echo.Context) error {
	userID := c.Get("userID").(int)

	stockSymbol := c.Param("stock_symbol")
	if stockSymbol == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid stock symbol"})
	}

	// İzleme listesinden hisseyi silme işlemi
	err := services.DeleteToWatchlist(userID, stockSymbol)
	if err != nil {
		if err.Error() == "stock not found in watchlist" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Stock not found in watchlist"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to remove stock from watchlist"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Stock removed from watchlist successfully"})

}
