package models

import "github.com/jinzhu/gorm"

// Note contains a single lecture note
type Note struct {
	gorm.Model

	Title       string
	Description string
	Body        string

	AuthorID uint
	Author   User

	Published bool

	Tags []Tag `gorm:"many2many:note_tag;"`

	// Ranking Info
	Ranking       int
	UpvoteUsers   []User `gorm:"many2many:upvotenote_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotenote_user;"`
}
