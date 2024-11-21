package services

import (
	"errors"
	"pulsefin/database"
	"pulsefin/models"
	"pulsefin/utils"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	query := "SELECT * FROM users"
	err := database.DB.Select(&users, query)

	return users, err
}

func UpdateEmail(userID int, password, newEmail string) error {
	// Kullanıcıyı ID ile bul
	var user models.User
	query := "SELECT id, password_hash, email FROM users WHERE id = $1"
	err := database.DB.Get(&user, query, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Şifreyi doğrula
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return errors.New("incorrect password")
	}

	// E-posta adresinin mevcut olup olmadığını kontrol et
	var existingUser models.User
	query = "SELECT id FROM users WHERE email = $1"
	err = database.DB.Get(&existingUser, query, newEmail)
	if err == nil {
		return errors.New("email already in use")
	}

	// Veritabanında e-posta adresini güncelle
	query = "UPDATE users SET email = $1 WHERE id = $2"
	_, err = database.DB.Exec(query, newEmail, user.ID)
	if err != nil {
		return errors.New("failed to update email")
	}

	return nil
}
