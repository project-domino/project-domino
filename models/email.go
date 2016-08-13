package models

import "github.com/jinzhu/gorm"

// An Email holds a message to be delivered to a user by email
type Email struct {
	gorm.Model

	UserID uint
	User   User

	Subject string
	Body    string

	Sent    bool
	Dropped bool
}

// MarkSent marks an email as sent, It also sets the dropped value to false
func (e *Email) MarkSent(db *gorm.DB) error {
	e.Sent = true
	e.Dropped = false

	err := db.Save(&e).Error

	return err
}

// MarkDropped marks an email as dropped, It also sets the sent value to false
func (e *Email) MarkDropped(db *gorm.DB) error {
	e.Sent = false
	e.Dropped = true

	err := db.Save(&e).Error

	return err
}
