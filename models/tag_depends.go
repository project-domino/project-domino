package models

import "github.com/jinzhu/gorm"

// TagDepends stores a tag - tag relation
// "Tag" depends on "Depends"
type TagDepends struct {
	gorm.Model

	Tag   Tag
	TagID uint

	Depends   Tag
	DependsID uint
}
