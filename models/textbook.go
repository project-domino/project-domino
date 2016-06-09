package models

import "github.com/jinzhu/gorm"

// An Textbook contains a set of collections
type Textbook struct {
	gorm.Model

	Title    string
	Intro    string
	Featured bool

	AuthorID uint
	Author   User

	Collections []Collection `gorm:"many2many:textbook_collection;"`
	Tags        []Tag        `gorm:"many2many:textbook_tag;"`
}
