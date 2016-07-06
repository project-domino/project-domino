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
func EditNote(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

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

// EditCollection returns the page to edit a given collection
func EditCollection(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	collectionID := c.Param("collectionID")

	// Query db for collection
	var collection models.Collection
	db.DB.Preload("Author").
		Preload("Tags").
		Preload("Notes").
		Where("id = ?", collectionID).First(&collection)
	if collection.ID == 0 {
		errors.CollectionNotFound.Apply(c)
		return
	}

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		errors.NotNoteOwner.Apply(c)
		return
	}

	// Format collection in JSON
	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		c.Error(err)
		errors.JSON.Apply(c)
		return
	}

	// Set request context and render html
	c.Set("user", user)
	c.Set("collection", collection)
	c.Set("collectionJSON", string(collectionJSON))
	c.HTML(200, "edit-collection.html", vars.Default(c))
}
