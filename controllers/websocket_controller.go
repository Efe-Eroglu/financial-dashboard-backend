package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

func StartWebSocketForUser(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok || userID <= 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing token"})
	}

	err := services.StartTickerWebSocketForUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start WebSocket connections"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "WebSocket connections started for user"})
}
