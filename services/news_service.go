package services

import (
	"pulsefin/database"
	"pulsefin/models"
)

func GetNews() ([]models.News, error) {
	var news []models.News

	query := "SELECT * FROM news"
	err := database.DB.Select(&news, query)

	return news, err
}
