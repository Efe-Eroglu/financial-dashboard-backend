package routes

import (
	"pulsefin/controllers"
	"pulsefin/middleware"

	"github.com/labstack/echo/v4"
)

func InitWebSocketRoutes(e *echo.Echo) {
	// WebSocket başlatma
	e.POST("/websocket/start", controllers.StartWebSocketForUser, middleware.AuthMiddleware)

	// WebSocket durdurma
	e.POST("/websocket/stop", controllers.StopWebSocketForUser, middleware.AuthMiddleware)
}
