package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// VerifyNotePublic checks if the note in the request context is public
func VerifyNotePublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire note from request context
		note := c.MustGet("note").(models.Note)

		// Check if the note is public
		if !note.Published {
			errors.NoteNotFound.Apply(c)
			return
		}

		c.Next()
	}
}
