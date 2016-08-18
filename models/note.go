package models

import "github.com/jinzhu/gorm"

// Note contains a single lecture note
type Note struct {
	gorm.Model

	Title       string
	Description string
	Body        string
	Featured    bool

	AuthorID uint
	Author   User

	Published bool

	Tags []Tag `gorm:"many2many:note_tag;"`

	// Ranking Info
	Ranking       int
	UpvoteUsers   []User `gorm:"many2many:upvotenote_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotenote_user;"`
}

// InList checks if a given note is in a list
func (n Note) InList(l []Note) bool {
	for _, e := range l {
		if n.ID == e.ID {
			return true
		}
	}
	return false
}
