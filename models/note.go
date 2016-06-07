package models

import "github.com/jinzhu/gorm"

// Note contains a single lecture note
// TODO escape html, verify string size < 2.5 mega characters
type Note struct {
	gorm.Model

	Title string
	Body  string
	URLID string

	WriterID uint
	Writer   User

	UniversityID uint
	University   University

	SubjectID uint
	Subject   Subject

	Outlines []Outline `gorm:"many2many:note_outline;"`
}
