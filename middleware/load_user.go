package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadUser loads a user into the request context
func LoadUser(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire username from URL
		username := c.Param("username")

		// Set objects to be preloaded to db
		preloadedDB := db.DB.Where("user_name = ?", username)
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// Query for user and set context
		var user models.User
		if err := preloadedDB.First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors.UserNotFound.Apply(c)
			} else {
				errors.DB.Apply(c)
			}
			return
		}

		c.Set("pageUser", user)

		c.Next()
	}
}
