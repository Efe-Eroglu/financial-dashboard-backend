package routes

import (
	"pulsefin/controllers"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func initStockRoutes(e *echo.Echo, db *sqlx.DB) {

	e.GET("/stocks", controllers.GetStocks(db))

}
