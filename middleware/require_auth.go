package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// RequireAuth is a middleware that gives the user an error if they are not
// logged in.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := c.MustGet("loggedIn").(bool)
		if !loggedIn {
			panic(errors.New("You must be logged in to perform this action."))
		}
		c.Next()
	}
}
