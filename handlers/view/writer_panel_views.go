package view

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/handlers/vars"
	"github.com/project-domino/project-domino/models"
)

// WriterPanelRedirect redirects the user to a page in the writer panel
// TODO redirect to latest draft
func WriterPanelRedirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/writer-panel/note")
}

// EditNote returns the page to edit a given note
// TODO This got merged, check for correctness.
func EditNote(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

	// Load Notes into the user
	db.DB.Model(&user).Association("Notes").Find(&user.Notes)

	// Query db for note
	var note models.Note
	db.DB.Preload("Author").
		Preload("Tags").
		Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		errors.NoteNotFound.Apply(c)
		return
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		errors.NotNoteOwner.Apply(c)
		return
	}

	// Format note in JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		c.Error(err)
		errors.JSON.Apply(c)
		return
	}

	// Set request context and render html
	c.Set("user", user)
	c.Set("note", note)
	c.Set("noteJSON", string(noteJSON))
	c.HTML(200, "edit-note.html", vars.Default(c))
}
