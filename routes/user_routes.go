package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func initUserRoutes(e *echo.Echo) {

	e.GET("/users", controllers.GetUsers)
}
