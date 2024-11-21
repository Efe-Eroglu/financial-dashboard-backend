package utils

import (
	"fmt"
	"net/smtp"
)

// SendEmail sends an email using the provided SMTP server
func SendEmail(to, subject, body string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "eferoglu1967@gmail.com"
	senderPassword := "bfkx fpxl beee wxcl"

	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
