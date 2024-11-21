package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

// GetNews godoc
// @Summary Haberleri Listele
// @Description Veritabanındaki tüm haberleri listeler
// @Tags News
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {object} map[string]string
// @Router /news [get]
func GetNews(c echo.Context) error {
	news, err := services.GetNews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
	}

	return c.JSON(http.StatusOK, news)
}
