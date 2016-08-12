package email

import (
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

var emails = make(chan models.Email)
var apiKey string

// Init initializes the email package
func Init(key string) error {
	apiKey = key

	go worker(emails)

	var err error

	// Get emails to be sent from the database
	var dbMails []models.Email
	err = db.DB.
		Where("sent = ?", false).
		Where("dropped = ?", false).
		Find(&dbMails).
		Error

	// Add emails to queue
	for _, e := range dbMails {
		emails <- e
	}

	return err
}
