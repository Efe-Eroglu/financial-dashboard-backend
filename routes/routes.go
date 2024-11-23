package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sqlx.DB) {
	initUserRoutes(e)
	initStockRoutes(e)
	initWatchlistRoutes(e)
	initNewsRoutes(e)
	initAuthRoutes(e)
	InitWebSocketRoutes(e)
}
