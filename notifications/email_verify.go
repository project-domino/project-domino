package notifications

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// EmailVerify saves a notification for a user to verify an email
func EmailVerify(db *gorm.DB, subject models.User) error {
	if subject.ID == 0 {
		return errors.InternalError
	}

	if err := RemoveEmailVerifyNotifications(db, subject); err != nil {
		return err
	}

	// Create new email verify notification
	if err := db.Create(&models.Notification{
		SubjectID: subject.ID,
		Type:      EmailVerifyNotificationType,
		Title:     "You must verify your email address.",
		Link:      "/email/verify",
	}).Error; err != nil {
		return err
	}

	return nil
}

// RemoveEmailVerifyNotifications marks all email verification notifications for
// a given user as read
func RemoveEmailVerifyNotifications(db *gorm.DB, subject models.User) error {
	err := db.Table("notifications").
		Where("subject_id = ?", subject.ID).
		Where("type = ?", EmailVerifyNotificationType).
		Updates(map[string]interface{}{"read": true}).
		Error

	return err
}
