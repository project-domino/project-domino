package models

import "github.com/jinzhu/gorm"

// An Tag describes an other object
type Tag struct {
	gorm.Model

	Name string

	Collections []Collection `gorm:"many2many:collection_tag;"`
	Notes       []Note       `gorm:"many2many:note_tag;"`
}
