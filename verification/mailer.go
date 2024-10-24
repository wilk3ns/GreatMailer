package verification

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"os"
)

func SendEmail(email string, hdr string, msg string) (string, error) {

	username := os.Getenv("GREATSTUFF_CONTACT_EMAIL")
	password := os.Getenv("GREATSTUFF_CONTACT_PASSWORD")

	smtpHost := "smtp.zoho.eu"
	smtpPort := 465

	message := msg

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", hdr)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err.Error(), err
	} else {
		fmt.Println("New e-mail from " + email)
		return "Email Sent Successfully!", nil
	}
}
