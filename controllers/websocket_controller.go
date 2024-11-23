package controllers

import (
	"net/http"
	"pulsefin/services"

	"github.com/labstack/echo/v4"
)

func StartWebSocket(c echo.Context) error {
	err := services.StartTickerWebSocket("BTC-USDT")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start WebSocket connection"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "WebSocket connection started"})
}
