package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

// GetWatchList godoc
// @Summary İzleme Listesini Listele
// @Description Veritabanındaki tüm izleme listesini listeler
// @Tags WatchList
// @Produce json
// @Success 200 {array} models.Watchlist
// @Failure 500 {object} map[string]string
// @Router /watchlist [get]
func GetWatchList(c echo.Context) error {
	watchList, err := services.GetWatchList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
	}

	return c.JSON(http.StatusOK, watchList)
}
