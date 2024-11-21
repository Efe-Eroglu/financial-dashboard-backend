package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func initStockRoutes(e *echo.Echo) {

	e.GET("/stocks", controllers.GetStocks)
	e.GET("/stocks/:stock_symbol", controllers.GetStock)

}
