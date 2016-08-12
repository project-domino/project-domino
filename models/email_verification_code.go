package models

import "github.com/jinzhu/gorm"

// An EmailVerificationCode holds the verification code for an individual
// user's email
type EmailVerificationCode struct {
	gorm.Model

	UserID uint
	User   User

	Email string

	Code string
}
