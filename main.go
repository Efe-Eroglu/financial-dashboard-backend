package main

import (
	"log"
	"net/http"
	"os"
	"pulsefin/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Dosya bulunamadı (.env)")
	}

	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "Server is running",
		})
	})

	port := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start((":" + port)))

}
