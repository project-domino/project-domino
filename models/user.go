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
	Notes        []Note       `gorm:"ForeignKey:AuthorID"`
	Collections  []Collection `gorm:"ForeignKey:AuthorID"`
	Textbooks    []Textbook   `gorm:"ForeignKey:AuthorID"`

	Email          string
	EmailVerified  bool
	SendNewsletter bool

	// Favorite Info
	FavoriteNotes       []Note       `gorm:"many2many:favoritenote_user;"`
	FavoriteCollections []Collection `gorm:"many2many:favoritecollection_user;"`

	// Ranking Info
	UpvoteCollections   []Collection `gorm:"many2many:upvotecollection_user;"`
	DownvoteCollections []Collection `gorm:"many2many:downvotecollection_user;"`

	UpvoteComments   []Comment `gorm:"many2many:upvotecomment_user;"`
	DownvoteComments []Comment `gorm:"many2many:downvotecomment_user;"`

	UpvoteNotes   []Note `gorm:"many2many:upvotenote_user;"`
	DownvoteNotes []Note `gorm:"many2many:downvotenote_user;"`

	UpvoteTextbooks   []Textbook `gorm:"many2many:upvotetextbook_user;"`
	DownvoteTextbooks []Textbook `gorm:"many2many:downvotetextbook_user;"`
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
