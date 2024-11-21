package routes

import (
	"pulsefin/controllers"
	"pulsefin/middleware"

	"github.com/labstack/echo/v4"
)

func initUserRoutes(e *echo.Echo) {

	e.GET("/users", controllers.GetUsers, middleware.AuthMiddleware)
}
