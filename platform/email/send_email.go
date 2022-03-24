package email

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

func SendEmailDestination(emailDestination, subject, message string) error {
	CONFIG_SENDER_NAME := os.Getenv("CONFIG_SENDER_NAME")
	CONFIG_AUTH_EMAIL := os.Getenv("CONFIG_AUTH_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("CONFIG_AUTH_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailDestination)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	// mailer.Attach("./logo.jpeg")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println("error send email ", err.Error())
		return err
	}

	log.Println("Mail sent!")
	return nil
}
