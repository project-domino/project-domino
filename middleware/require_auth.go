package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// RequireAuth is a middleware that gives the user an error if they are not
// logged in.
// TODO: Redirect the user to login page w/ redirect-back GET.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := c.MustGet("loggedIn").(bool)
		if !loggedIn {
			c.AbortWithError(401, errors.New("You must be logged in to perform this action."))
		}
		c.Next()
	}
}
