package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// VerifyNoteOwner verifies if the request user is the owner of the
// note in the request context
// A user and note must be in the request context
func VerifyNoteOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire variables
		user := c.MustGet("user").(models.User)
		note := c.MustGet("note").(models.Note)

		// Check if request user is the owner of the note
		if note.AuthorID != user.ID {
			errors.NotNoteOwner.Apply(c)
			return
		}

		c.Next()
	}
}
