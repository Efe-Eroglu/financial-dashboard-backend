package controllers

import (
	"fmt"
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

func ResetPassword(c echo.Context) error {
	var request struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	err := services.ResetPassword(request.Email, request.OldPassword, request.NewPassword)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		if err.Error() == "incorrect old password" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Incorrect old password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to reset password"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
}

func ForgotPassword(c echo.Context) error {
	var request struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&request); err != nil || request.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email format or missing email"})
	}

	resetCode, err := services.CreatePasswordResetRequest(request.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to process request"})
	}

	fmt.Printf("Mock email sent to %s: Reset code: %s\n", request.Email, resetCode)

	return c.JSON(http.StatusOK, map[string]string{"message": "If the email exists, a password reset code has been sent"})
}
