package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func initNewsRoutes(e *echo.Echo) {
	e.GET("/news", controllers.GetNews)
}
