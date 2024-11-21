package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

// GetStocks godoc
// @Summary Hisse Senetlerini Listele
// @Description Veritabanındaki tüm hisse senetlerini listeler
// @Tags Stocks
// @Produce json
// @Success 200 {array} models.Stock
// @Failure 500 {object} map[string]string
// @Router /stocks [get]
func GetStocks(c echo.Context) error {
	stocks, err := services.GetStocks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
	}

	return c.JSON(http.StatusOK, stocks)
}
