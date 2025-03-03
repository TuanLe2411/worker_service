package mail

import (
	"net/smtp"
	"os"
)

func SendEmail(to, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.Auth(smtp.PlainAuth("", email, password, smtpHost))
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, []byte(body))
}
