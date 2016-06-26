package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// RequireUserType is a middleware that gives the user an error if they are not
// of one of the given types.
func RequireUserType(types ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		for _, userType := range types {
			if userType == user.Type {
				c.Next()
				return
			}
		}

		errors.BadUserType.Apply(c)
	}
}
