package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadRequestUser loads certain objects in the request context's user
func LoadRequestUser(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire user from the request context.
		user := c.MustGet("user").(models.User)

		if user.ID != 0 {
			// Set objects to be preloaded to db
			preloadedDB := db.DB.Where("id = ?", user.ID)
			for _, object := range objects {
				preloadedDB = preloadedDB.Preload(object)
			}

			// Query for user and set context
			var loadedUser models.User
			if err := preloadedDB.First(&loadedUser).Error; err != nil {
				errors.DB.Apply(c)
				return
			}
			c.Set("user", loadedUser)
		}

		c.Next()
	}
}
