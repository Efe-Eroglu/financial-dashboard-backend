package main

import (
	"log"
	"pulsefin/config"
	"pulsefin/database"
	"pulsefin/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()

	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.InitRoutes(e, db)

	port := config.AppConfig.ServerPort
	log.Printf("Server is running on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
