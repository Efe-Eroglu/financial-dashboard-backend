package services

import (
	"pulsefin/database"
	"pulsefin/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	query := "SELECT * FROM users"
	err := database.DB.Select(&users, query)

	return users, err
}
