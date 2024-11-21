package models

import "time"

type PasswordReset struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	ResetCode string    `json:"reset_code" db:"reset_code"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	Used      bool      `json:"used" db:"used"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
