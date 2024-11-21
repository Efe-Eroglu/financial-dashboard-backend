package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"pulsefin/database"
	"pulsefin/models"
	"pulsefin/utils"
	"time"
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

func GenerateResetCode() string {
	b := make([]byte, 3)
	rand.Read(b)
	return fmt.Sprintf("%06d", (int(b[0])<<16+int(b[1])<<8+int(b[2]))%1000000)
}

func CreatePasswordResetRequest(email string) (string, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	err := database.DB.Get(&count, query, email)
	if err != nil || count == 0 {
		return "", errors.New("user not found")
	}

	resetCode := GenerateResetCode()
	expiresAt := time.Now().Add(15 * time.Minute)

	query = "INSERT INTO password_resets (email, reset_code, expires_at) VALUES ($1, $2, $3)"
	_, err = database.DB.Exec(query, email, resetCode, expiresAt)
	if err != nil {
		return "", errors.New("failed to create password reset request")
	}
	sendMockEmail(email, resetCode)

	return resetCode, nil
}

func sendMockEmail(email, resetCode string) {
	fmt.Printf("Mock Email Sent to %s: Your reset code is %s\n", email, resetCode)
}
