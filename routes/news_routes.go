package routes

import (
	"pulsefin/controllers"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func initNewsRoutes(e *echo.Echo, db *sqlx.DB) {
	e.GET("/news", controllers.GetNews(db))
}
