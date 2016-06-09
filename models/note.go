package models

import "github.com/jinzhu/gorm"

// Note contains a single lecture note
type Note struct {
	gorm.Model

	Title string
	Body  string

	AuthorID uint
	Author   User

	Collections []Collection `gorm:"many2many:note_collection;"`
	Tags        []Tag        `gorm:"many2many:note_tag;"`
}
