package models

import "github.com/jinzhu/gorm"

// University holds the values for an individual university.
// ShortName must be unique
type University struct {
	gorm.Model

	Name      string
	ShortName string
}
