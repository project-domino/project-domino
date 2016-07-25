package models

import "github.com/jinzhu/gorm"

// An Email holds a message to be delivered to a user by email
type Email struct {
	gorm.Model

	UserID uint
	User   User

	Subject string
	Body    string

	// This should be set to true when sending a verification email
	Verification bool
}
