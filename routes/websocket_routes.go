package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func InitWebSocketRoutes(e *echo.Echo) {
	e.POST("/websocket/start", controllers.StartWebSocket)
}
