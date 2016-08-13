package models

import "github.com/jinzhu/gorm"

// A Notification holds a user notification
type Notification struct {
	gorm.Model

	ActorID   uint
	Actor     User
	SubjectID uint
	Subject   User

	Type    string
	Title   string
	Message string
	Link    string
	Read    bool
}
