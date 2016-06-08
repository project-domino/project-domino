package models

import "github.com/jinzhu/gorm"

// Note contains a single lecture note
// TODO escape html, verify string size < 2.5 mega characters
type Note struct {
	gorm.Model

	Title string
	Body  string

	WriterID uint
	Writer   User

	UniversityID uint
	University   University

	Collections []Collection `gorm:"many2many:note_collection;"`
	Tags        []Tag        `gorm:"many2many:note_tag;"`
}
