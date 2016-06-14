package models

import "github.com/jinzhu/gorm"

// NoteOrder contains a single note (prev, next) relationship
type NoteOrder struct {
	gorm.Model

	Prev   Note
	PrevID uint
	Next   Note
	NextID uint

	Weight int
}
