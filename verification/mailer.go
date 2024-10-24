package verification

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(email string, hdr string, msg string) (string, error) {

	username := "info@greatstuff.ee"
	password := "bta3stW&"

	smtpHost := "smtp.zoho.eu"
	smtpPort := 465

	message := msg

	m := gomail.NewMessage()

	m.SetHeader("From", email)

	m.SetHeader("To", "info@greatstuff.ee")

	m.SetHeader("Subject", hdr)

	m.SetBody("text/plain", message)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err.Error(), err
	} else {
		fmt.Println("Email Sent to " + email)
		return "Email Sent Successfully!", nil
	}
}
