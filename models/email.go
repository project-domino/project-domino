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
