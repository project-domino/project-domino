package models

import "github.com/jinzhu/gorm"

// An Collection contains a set of notes
type Collection struct {
	gorm.Model

	Title       string
	Description string
	Featured    bool

	AuthorID uint
	Author   User

	Published bool

	Tags  []Tag  `gorm:"many2many:collection_tag;"`
	Notes []Note `gorm:"-"`

	// Ranking Info
	Ranking       int
	UpvoteUsers   []User `gorm:"many2many:upvotecollection_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotecollection_user;"`
}

// InList checks if a given collection is in a list
func (c Collection) InList(l []Collection) bool {
	for _, e := range l {
		if c.ID == e.ID {
			return true
		}
	}
	return false
}
