package models

import "github.com/jinzhu/gorm"

// A Subject contains the information for a single subject
// Name must be unique
type Subject struct {
	gorm.Model

	Name        string
	Description string
	Topics      []Outline `gorm:"many2many:subject_topic;"`
}
