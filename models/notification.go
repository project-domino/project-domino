package models

import "github.com/jinzhu/gorm"

// These constants are the valid values for Notification.Type.
const (
	CommentNotificationType string = "comment"
)

// A Notification holds a user notification
type Notification struct {
	gorm.Model

	SendorID uint
	Sendor User
	UserID uint
	User   User

	Title string
	Type string
	Payload string
	Message string
	Read bool
}
