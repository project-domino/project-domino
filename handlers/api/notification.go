package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// MarkNotificationRead marks a notification as read
func MarkNotificationRead(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	notificationID := c.Param("notificationID")

	var notification models.Notification
	if err := db.DB.
		Where("id = ?", notificationID).
		Where("subject_id = ?", user.ID).
		First(&notification).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			errors.NotificationNotFound.Apply(c)
		} else {
			errors.DB.Apply(c)
		}

		return
	}

	notification.Read = true

	if err := db.DB.Save(&notification).Error; err != nil {
		errors.DB.Apply(c)
		return
	}
}
