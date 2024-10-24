package verification

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"os"
	"strconv"
)

func SendEmail(email string, hdr string, msg string) (string, error) {

	username := os.Getenv("GREATSTUFF_CONTACT_EMAIL")
	password := os.Getenv("GREATSTUFF_CONTACT_PASSWORD")

	smtpHost := os.Getenv("GREATSTUFF_EMAIL_OUTGOING")
	smtpPort, _ := strconv.Atoi(os.Getenv("GREATSTUFF_EMAIL_PORT"))

	message := msg

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", username)
	m.SetHeader("Subject", hdr)
	m.SetBody("text/plain", "Message from:"+" "+email+"\n"+message)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("failed to send email " + email)
		return err.Error(), err
	} else {
		fmt.Println("New e-mail from " + email)
		return "Email Sent Successfully!", nil
	}
}
