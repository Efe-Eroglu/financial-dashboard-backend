package services

import (
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

func CreatePasswordResetRequest(email string) (string, error) {
	// Kullanıcı var mı kontrol et
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	err := database.DB.Get(&count, query, email)
	if err != nil || count == 0 {
		fmt.Println("Error: User not found or database query failed:", err)
		return "", errors.New("user not found")
	}

	// Reset kodu oluştur
	resetCode, err := utils.GenerateStrongResetCode(12)
	if err != nil {
		fmt.Println("Error: Failed to generate reset code:", err)
		return "", errors.New("failed to generate reset code")
	}

	// Veritabanına reset kodunu ekle
	expiresAt := time.Now().Add(15 * time.Minute)
	query = "INSERT INTO password_resets (email, reset_code, expires_at) VALUES ($1, $2, $3)"
	_, err = database.DB.Exec(query, email, resetCode, expiresAt)
	if err != nil {
		fmt.Println("Error: Failed to insert reset code into database:", err)
		return "", errors.New("failed to create password reset request")
	}

	// E-posta gönder
	subject := "Password Reset Request"
	body := fmt.Sprintf("Your password reset code is: %s\nThis code is valid for 15 minutes.", resetCode)
	err = utils.SendEmail(email, subject, body)
	if err != nil {
		fmt.Println("Error: Failed to send email:", err)
		return "", errors.New("failed to send reset email")
	}

	return resetCode, nil
}

func ResetPasswordWithCode(email, resetCode, newPassword string) error {
	// Reset kodunun doğruluğunu kontrol et
	var count int
	query := "SELECT COUNT(*) FROM password_resets WHERE email = $1 AND reset_code = $2 AND expires_at > NOW() AND used = FALSE"
	err := database.DB.Get(&count, query, email, resetCode)
	if err != nil || count == 0 {
		return errors.New("invalid or expired reset code")
	}

	// Yeni şifreyi hashle
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Kullanıcının şifresini güncelle
	query = "UPDATE users SET password_hash = $1 WHERE email = $2"
	_, err = database.DB.Exec(query, hashedPassword, email)
	if err != nil {
		return errors.New("failed to update password")
	}

	// Reset kodunu kullanılmış olarak işaretle
	query = "UPDATE password_resets SET used = TRUE WHERE email = $1 AND reset_code = $2"
	_, err = database.DB.Exec(query, email, resetCode)
	if err != nil {
		return errors.New("failed to mark reset code as used")
	}

	return nil
}
