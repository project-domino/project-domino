package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// Login adds a user struct to the request context based on the authentication
// token provided in a cookie. Also sets a loggedIn boolean.
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if auth cookie is present.
		authCookie, err := c.Cookie("auth")

		// Go to next handler if no cookie is found
		// Set loggedIn to false
		if err == http.ErrNoCookie {
			notLoggedIn(c)
			return
		} else if err != nil {
			c.AbortWithError(500, err)
			return
		}

		// If the cookie is present, search the database for the token.
		var authEntries []models.AuthToken

		if err := db.DB.Limit(1).
			Preload("User").
			Where("token = ?", authCookie).
			Where("expires > ?", time.Now()).
			Find(&authEntries).
			Error; err != nil && err != gorm.ErrRecordNotFound {
			errors.DB.Apply(c)
			return
		}
		if len(authEntries) == 0 {
			// Clear the invalid/expired authtoken.
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "auth",
				Path:    "/",
				Value:   "",
				Expires: time.Unix(0, 0),
			})
			notLoggedIn(c)
			return
		}

		// If there is a token and it is not expired, add the user to the context.
		auth := authEntries[0]
		c.Set("user", auth.User)
		c.Set("loggedIn", true)

		c.Next()
	}
}

func notLoggedIn(c *gin.Context) {
	c.Set("user", models.User{})
	c.Set("loggedIn", false)
	c.Next()
}
