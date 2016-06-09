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

	Notes     []Note     `gorm:"many2many:note_collection;"`
	Textbooks []Textbook `gorm:"many2many:textbook_collection;"`
	Tags      []Tag      `gorm:"many2many:collection_tag;"`
}
