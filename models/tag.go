package models

import "github.com/jinzhu/gorm"

// An Tag describes an other object
type Tag struct {
	gorm.Model

	Name      string
	DependsOn []Tag `gorm:"foreignkey:tag_id;associationforeignkey:depends_id;many2many:tag_depends;"`

	Collections []Collection `gorm:"many2many:collection_tag;"`
	Notes       []Note       `gorm:"many2many:note_tag;"`
	Textbooks   []Textbook   `gorm:"many2many:textbook_tag;"`
}
