package main

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"

	"github.com/project-domino/project-domino/models"
)

// SendEmail sends the passed email
func SendEmail(e models.Email, auth smtp.Auth, addr string) error {
	// Verify email is valid
	// TODO verify more
	if e.User.Email == "" {
		return errors.New("Email not valid.")
	}

	// If this email is not for verification, check if the user's email is verified
	if !(e.Verification && e.User.EmailVerified) {
		return errors.New("Email Address not verified")
	}

	// Create email body
	body := []byte(
		fmt.Sprintf("To: %s\r\n", e.User.Email) +
			fmt.Sprintf("Subject: %s\r\n", e.Subject) +
			"\r\n" +
			fmt.Sprintf("%s \r\n", e.Body))

	// Send email
	err := smtp.SendMail(
		addr,
		auth,
		"no-reply@notebox.org",
		[]string{e.User.Email},
		body,
	)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
