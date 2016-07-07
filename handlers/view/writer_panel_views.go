package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func EditNote(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	note := c.MustGet("note").(models.Note)

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		errors.NotNoteOwner.Apply(c)
		return
	}

	// Render HTML
	c.HTML(200, "edit-note.html", vars.Default(c))
}

// EditCollection returns the page to edit a given collection
func EditCollection(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	collection := c.MustGet("collection").(models.Collection)

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		errors.NotNoteOwner.Apply(c)
		return
	}

	// Render HTML
	c.HTML(200, "edit-collection.html", vars.Default(c))
}
