package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func initAuthRoutes(e *echo.Echo) {
	e.POST("/auth/register", controllers.RegisterUser)
	e.POST("/auth/login", controllers.LoginUser)
}