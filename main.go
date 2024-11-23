package main

import (
	"log"
	"pulsefin/config"
	"pulsefin/database"
	"pulsefin/routes"

	_ "pulsefin/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()

	database.ConnectDB()
	defer database.DB.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.InitRoutes(e, database.DB)

	port := config.AppConfig.ServerPort
	log.Printf("Server is running on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
