package routes

import (
	"pulsefin/controllers"
	"pulsefin/middleware"

	"github.com/labstack/echo/v4"
)

func initAuthRoutes(e *echo.Echo) {
	e.POST("/auth/register", controllers.RegisterUser)
	e.POST("/auth/login", controllers.LoginUser)
	e.PUT("/auth/reset-password", controllers.ResetPassword, middleware.AuthMiddleware)
	e.POST("/auth/forgot-password", controllers.ForgotPassword)
	e.POST("/auth/reset-password-code", controllers.ResetPasswordWithCode)
}
