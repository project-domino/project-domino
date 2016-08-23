package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// These constants are the valid values by in LoadNotes
const (
	LoadNotesAuthor      string = "author"
	LoadNotesRequestUser        = "requestUser"
	LoadNotesUserUpvote         = "requestUserUpvote"
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

		preloadedDB = preloadedDB.
			Limit(items).
			Offset((page - 1) * items).
			Order("updated_at desc")

		var notes []models.Note
		var err error

		switch by {
		case LoadNotesAuthor:
			user := c.MustGet("pageUser").(models.User)
			err = preloadedDB.
				Where("author_id = ?", user.ID).
				Where("published = ?", true).
				Find(&notes).
				Error
		case LoadNotesUserUpvote:
			user := c.MustGet("pageUser").(models.User)
			err = preloadedDB.
				Model(&user).
				Association("UpvoteNotes").
				Find(&notes).
				Error
		case LoadNotesRequestUser:
			user := c.MustGet("user").(models.User)
			err = preloadedDB.
				Where("author_id = ?", user.ID).
				Find(&notes).
				Error
		default:
			errors.InternalError.Apply(c)
			return
		}

		if err != nil {
			errors.DB.Apply(c)
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
