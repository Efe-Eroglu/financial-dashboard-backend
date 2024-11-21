package controllers

import (
	"net/http"
	"pulsefin/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetWatchList(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var watch_list []models.Watchlist

		query := "SELECT * FROM watchlist"
		err := db.Select(&watch_list, query)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		}

		return c.JSON(http.StatusOK, watch_list)
	}
}
