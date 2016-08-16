package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/notifications"
)

// LoadNotifications loads unread notifications into the user
func LoadNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)

		var n []models.Notification
		db.DB.
			Order(fmt.Sprintf("read, type = '%s', created_at desc", notifications.EmailVerifyNotificationType)).
			Where("subject_id = ?", user.ID).
			Limit(10).
			Find(&n)

		c.Set("notifications", n)

		c.Next()
	}
}
