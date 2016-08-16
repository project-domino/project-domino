package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// LoadNotifications loads unread notifications into the user
func LoadNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)

		var notifications []models.Notification
		db.DB.Model(&user).Where("read = ?", false).Association("Notifications").Find(&notifications)

		c.Set("notifications", notifications)

		c.Next()
	}
}
