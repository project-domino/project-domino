package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// LoadUser loads certain objects in the request context's user
func LoadUser(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire user from the request context.
		user := c.MustGet("user").(models.User)

		// Set objects to be preloaded to db
		preloadedDB := util.DB.Where("id = ?", user.ID)
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
