package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// VerifyCollectionOwner verifies if the request user is the owner of the
// collection in the request context
// A user and collection must be in the request context
func VerifyCollectionOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire variables
		user := c.MustGet("user").(models.User)
		collection := c.MustGet("collection").(models.Collection)

		// Check if request user is the owner of the collection
		if collection.AuthorID != user.ID {
			errors.NotNoteOwner.Apply(c)
			return
		}

		c.Next()
	}
}
