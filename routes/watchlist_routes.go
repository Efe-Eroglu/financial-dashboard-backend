package routes

import (
	"pulsefin/controllers"

	"github.com/labstack/echo/v4"
)

func initWatchlistRoutes(e *echo.Echo) {

	e.GET("watch-lists", controllers.GetWatchList)

}
