package controllers

import (
	"net/http"
	"pulsefin/models"
	"pulsefin/services"
	"pulsefin/utils"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Inpput"})
	}

	err := services.RegisterUser(&user)
	if err != nil {
		if err.Error() == "email already in use" {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already in use"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func LoginUser(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
