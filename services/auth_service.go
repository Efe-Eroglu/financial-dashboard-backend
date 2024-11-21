package services

import (
	"errors"
	"pulsefin/database"
	"pulsefin/models"
	"pulsefin/utils"
)

func RegisterUser(user *models.User) error {
	existingUser := models.User{}
	query := "SELECT id FROM users WHERE email = $1"
	err := database.DB.Get(&existingUser, query, user.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword

	query = "INSERT INTO users (username, email, password_hash, created_at) VALUES ($1, $2, $3, NOW())"
	_, err = database.DB.Exec(query, user.Username, user.Email, user.PasswordHash)
	return err
}

func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, email, password_hash FROM users WHERE email = $1"
	err := database.DB.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func ResetPassword(email, oldPassword, newPassword string) error {
	var user models.User
	query := "SELECT id, password_hash FROM users WHERE email = $1"
	err := database.DB.Get(&user, query, email)
	if err != nil {
		return errors.New("user not found")
	}

	if !utils.CheckPasswordHash(oldPassword, user.PasswordHash) {
		return errors.New("incorrect old password")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	query = "UPDATE users SET password_hash = $1 WHERE id = $2"
	_, err = database.DB.Exec(query, hashedPassword, user.ID)
	if err != nil {
		return errors.New("failed to update password")
	}

	return nil
}
