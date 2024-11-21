package utils

import (
	"fmt"
	"net/smtp"
	"pulsefin/config"
)

func SendEmail(to, subject, body string) error {
	smtpHost := config.AppConfig.SMTPHost
	smtpPort := config.AppConfig.SMTPPort
	senderEmail := config.AppConfig.SMTPEmail
	senderPassword := config.AppConfig.SMTPPassword

	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
