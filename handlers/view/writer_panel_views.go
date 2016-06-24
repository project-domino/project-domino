package view

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/handlers/vars"
	"github.com/project-domino/project-domino/models"
)

// EditNote returns the page to edit a given note
func EditNote(c *gin.Context) {
	// Acquire variables from request context.
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

	// Query db for note
	var note models.Note
	db.Preload("Author").
		Preload("Tags").
		Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		c.AbortWithError(404, errors.New("Note not found"))
		return
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		c.AbortWithError(403, errors.New("You are not the owner of this note"))
		return
	}

	// Format note in JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		c.AbortWithError(500, errors.New("Could not convert note to json"))
		return
	}

	// Set request context and render html
	c.Set("user", user)
	c.Set("note", note)
	c.Set("noteJSON", string(noteJSON))
	c.HTML(200, "edit-note.html", vars.Default(c))
}
