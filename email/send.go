package email

import (
	"errors"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Send queues an email to be sent
// The email is also added to the db
func Send(e models.Email) error {
	// Check if email to be sent has a user
	if e.User.ID == 0 {
		return errors.New("Email requires a user.")
	}

	// Check if the user has an email
	if e.User.Email == "" {
		return errors.New("User does not have an email.")
	}

	// Set sent and dropped values to false
	e.Sent = false
	e.Dropped = false

	// Save email in db
	if err := db.DB.Create(&e).Error; err != nil {
		return err
	}

	// Add email to send queue
	emails <- e

	return nil
}
