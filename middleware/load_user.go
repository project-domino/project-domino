package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// LoadUser loads certain objects in the request context's user
func LoadUser(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire variables from request context.
		db := c.MustGet("db").(*gorm.DB)
		user := c.MustGet("user").(models.User)

		// Set objects to be preloaded to db
		preloadedDB := db.Where("id = ?", user.ID)
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// Query for user and set context
		var loadedUser models.User
		preloadedDB.First(&loadedUser)
		c.Set("user", loadedUser)

		c.Next()
	}
}
