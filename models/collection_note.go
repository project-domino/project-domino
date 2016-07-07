package models

import "github.com/jinzhu/gorm"

// CollectionNote holds the relationship between a note and a collection
// Order begins at 1, this represents the first note in the collection
type CollectionNote struct {
	gorm.Model

	Note   Note
	NoteID uint

	Collection   Collection
	CollectionID uint `gorm:"index"`

	Order uint
}
