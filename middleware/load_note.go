package middleware

import (
	"encoding/json"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadNote loads a note into the request context with specified objects
// It also loads noteJSON, the note object serialized into JSON
// :noteID must be in the URL
func LoadNote(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire noteID from URL
		noteID := c.Param("noteID")

		// Set objects to be preloaded to db
		preloadedDB := db.DB.Where("id = ?", noteID)
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// Query for note
		var note models.Note
		preloadedDB.First(&note)
		if note.ID == 0 {
			errors.NoteNotFound.Apply(c)
			return
		}

		// Format note in JSON
		noteJSON, err := json.Marshal(note)
		if err != nil {
			c.Error(err)
			errors.JSON.Apply(c)
			return
		}

		// Add note and noteJSON to request context
		c.Set("note", note)
		c.Set("noteJSON", string(noteJSON))
		// TODO remove only scripts
		c.Set("noteHTML", template.HTML(note.Body))

		c.Next()
	}
}
