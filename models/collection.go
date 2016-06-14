package models

import "github.com/jinzhu/gorm"

// An Collection contains a set of notes
type Collection struct {
	gorm.Model

	Title    string
	Intro    string
	Featured bool

	AuthorID uint
	Author   User

	Notes []Note `gorm:"many2many:note_collection;"`
	Tags  []Tag  `gorm:"many2many:collection_tag;"`

	// Ranking Info
	Upvotes       uint
	Downvotes     uint
	UpvoteUsers   []User `gorm:"many2many:upvotecollection_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotecollection_user;"`
}
