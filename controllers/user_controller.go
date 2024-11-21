package controllers

import (
	"net/http"
	"pulsefin/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetUsers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User

		query := "SELECT * FROM users"
		err := db.Select(&users, query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		}

		return c.JSON(http.StatusOK, users)
	}
}
