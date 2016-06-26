package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
)

// RequireAuth is a middleware that gives the user an error if they are not
// logged in.
// TODO: Redirect the user to login page w/ redirect-back GET.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn := c.MustGet("loggedIn").(bool)
		if !loggedIn {
			errors.LoginRequired.Apply(c)
			return
		}
		c.Next()
	}
}
