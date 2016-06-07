package models

import "github.com/jinzhu/gorm"

// These constants are the valid values for Outline.Type.
const (
	Textbook string = "textbook"
	List            = "list"
	Topic           = "topic"
)

// An Outline contains a set of notes
type Outline struct {
	gorm.Model

	Title       string
	Description string
	Type        string
	URLID       string
	IntroID     uint
	Intro       Note
	Subjects    []Subject `gorm:"many2many:subject_topic;"`
	Notes       []Note    `gorm:"many2many:note_outline;"`
}
