package routes

import (
	"pulsefin/controllers"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func initUserRoutes(e *echo.Echo, db *sqlx.DB) {

	e.GET("/users", controllers.GetUsers(db))
}
