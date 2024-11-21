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

func UpdateEmail(c echo.Context) error {
	userID := c.Get("userID").(int)

	var request struct {
		NewEmail string `json:"new_email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	err := services.UpdateEmail(userID, request.Password, request.NewEmail)
	if err != nil {
		if err.Error() == "incorrect password" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Incorrect password"})
		}
		if err.Error() == "email already in use" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already in use"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update email"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email updated successfully"})
}
