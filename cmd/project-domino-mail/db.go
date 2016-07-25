package main

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// GetEmail gets the next email to be sent from the server
func GetEmail(db *gorm.DB) (models.Email, error) {
	var email models.Email

	if err := db.
		Preload("User").
		Order("created_at").
		First(&email).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		return email, err
	}

	return email, nil
}

// MarkSent marks an email as sent
func MarkSent(db *gorm.DB, e models.Email) error {
	// Check if email has primary key
	// This check is necessary. An email with an ID of 0 passed to db.Delete will
	// delete all emails in the database
	if e.ID == 0 {
		return errors.New("Email must have ID.")
	}

	// Delete email
	// NOTE: This is a soft delete. Email will still be in the database with the
	// DeletedAt field set to the current time.
	if err := db.Delete(&e).Error; err != nil {
		return err
	}
	return nil
}
