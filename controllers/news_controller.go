package controllers

import (
	"net/http"
	"pulsefin/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetNews(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var news []models.News

		query := "SELECT * FROM news"
		err := db.Select(&news, query)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		}

		return c.JSON(http.StatusOK, news)
	}
}
