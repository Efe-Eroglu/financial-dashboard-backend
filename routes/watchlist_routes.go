package routes

import (
	"pulsefin/controllers"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func initWatchlistRoutes(e *echo.Echo, db *sqlx.DB) {

	e.GET("watch-lists", controllers.GetWatchList(db))

}
