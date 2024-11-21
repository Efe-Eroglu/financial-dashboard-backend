package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

// GetUsers godoc
// @Summary Kullanıcıları Listele
// @Description Veritabanındaki tüm kullanıcıları listeler
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
	}

	return c.JSON(http.StatusOK, users)
}
