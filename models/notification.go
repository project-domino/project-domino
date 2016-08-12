package models

import "github.com/jinzhu/gorm"

// These constants are valid values for Notification.type
const (
	emailVerifyNotificationType string = "email_verify"
)

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

// GetEmailVerifyNotification returns a notification to verify an email
func GetEmailVerifyNotification(subject User) Notification {
	return Notification{
		SubjectID: subject.ID,
		Type:      emailVerifyNotificationType,
		Title:     "You must verify your email address.",
		Link:      "/email/verify",
	}
}
