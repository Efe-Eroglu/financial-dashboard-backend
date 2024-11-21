package routes

import (
	"pulsefin/controllers"
	"pulsefin/middleware"

	"github.com/labstack/echo/v4"
)

func initWatchlistRoutes(e *echo.Echo) {

	e.GET("/watchlists", controllers.GetWatchList)
	e.POST("/watchlist", controllers.AddToWatchlist, middleware.AuthMiddleware)
	e.DELETE("/watchlist/:stock_symbol", controllers.DeleteToWatchlist, middleware.AuthMiddleware)

}
