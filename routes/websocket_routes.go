package routes

import (
	"pulsefin/controllers"
	"pulsefin/middleware"

	"github.com/labstack/echo/v4"
)

func InitWebSocketRoutes(e *echo.Echo) {
	e.POST("/websocket/start", controllers.StartWebSocketForUser, middleware.AuthMiddleware)
}
