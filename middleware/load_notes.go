package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// These constants are the valid values by in LoadNotes
const (
	LoadNotesAuthor string = "author"
)

// LoadNotes loads notes into the request context with specified objects
// It also loads notesJSON, the notes serialized into JSON
// The first argument must be what the notes are loaded by
// "by" can be one of the following options: author
func LoadNotes(by string, objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get list vars
		items, page, varErr := getListVars(c)
		c.Set("items", items)
		if varErr != nil {
			varErr.Apply(c)
			return
		}

		// Set objects to be preloaded to db
		preloadedDB := db.DB
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		switch by {
		case "author":
			// Acquire username from URL
			username := c.Param("username")
			// Find user in db
			var user models.User
			if err := db.DB.
				Where("user_name = ?", username).
				First(&user).
				Error; err != nil {

				if err == gorm.ErrRecordNotFound {
					errors.UserNotFound.Apply(c)
				} else {
					c.AbortWithError(500, err)
				}
				return
			}
			// add where statement to query
			preloadedDB = preloadedDB.Where("author_id = ?", user.ID)

		default:
			errors.InternalError.Apply(c)
			return
		}

		// Query for notes
		var notes []models.Note
		if err := preloadedDB.
			Limit(items).
			Offset((page - 1) * items).
			Find(&notes).
			Error; err != nil {

			c.AbortWithError(500, err)
			return
		}

		// Format notes in JSON
		notesJSON, err := json.Marshal(notes)
		if err != nil {
			errors.JSON.Apply(c)
			return
		}

		// Set pages
		prevPage, nextPage := getPages(page, len(notes) < items)
		c.Set("currentPage", page)
		c.Set("nextPage", nextPage)
		c.Set("prevPage", prevPage)

		// Add notes and notesJSON to request context
		c.Set("notes", notes)
		c.Set("notesJSON", string(notesJSON))

		c.Next()
	}
}
