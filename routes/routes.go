package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sqlx.DB) {
	initUserRoutes(e, db)
	initStockRoutes(e, db)
	initWatchlistRoutes(e, db)
	initNewsRoutes(e, db)
}
