package controllers

import (
	"net/http"
	"pulsefin/models"
	"pulsefin/services"
	"pulsefin/utils"

	"github.com/labstack/echo/v4"
)

// RegisterUser godoc
// @Summary Kullanıcı Kaydı
// @Description Yeni bir kullanıcı kaydı yapar
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Kullanıcı bilgileri"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
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

// LoginUser godoc
// @Summary Kullanıcı Girişi
// @Description Kullanıcı giriş yapar ve JWT döndürür
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body object true "Kullanıcı giriş bilgileri"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
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
