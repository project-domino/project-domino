package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// These constants are the valid values for User.Type.
const (
	Admin   string = "admin"
	Writer         = "writer"
	General        = "general"
)

// A User is a user of the website. They can either be a admin, writer, or general.
type User struct {
	gorm.Model

	Type      string
	FirstName string
	LastName  string
	UserName  string
	Passhash  string

	// Only for writer
	UniversityID uint
	University   University

	Email          string
	EmailVerified  bool
	SendNewsletter bool
}

// CheckPassword checks if the provided password is correct. Note that it will
// return false whether the password was incorrect or an error is encountered,
// with no means to disambiguate the two.
func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Passhash), []byte(password)) == nil
}

// SetPassword hashes the provided password with bcrypt with the default cost,
// currently 10.
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return err
	}
	u.Passhash = string(hash)
	return nil
}
