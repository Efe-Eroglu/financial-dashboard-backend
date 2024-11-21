package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

func GetStocks(c echo.Context) error {
	stocks, err := services.GetStocks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
	}

	return c.JSON(http.StatusOK, stocks)
}

func GetStock(c echo.Context) error {
	stockSymbol := c.Param("stock_symbol")
	if stockSymbol == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid stock symbol"})
	}

	stock, err := services.GetStock(stockSymbol)
	if err != nil {
		if err.Error() == "stock not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Stock not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch stock"})
	}

	return c.JSON(http.StatusOK, stock)
}
