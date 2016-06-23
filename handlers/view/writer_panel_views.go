package view

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/handlers/vars"
	"github.com/project-domino/project-domino/models"
)

// WriterPanelSimple returns a handler for a specified writer panel page
func WriterPanelSimple(template string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire variables from request context.
		db := c.MustGet("db").(*gorm.DB)
		user := c.MustGet("user").(models.User)

		// Load Notes into the user
		db.Model(&user).Association("Notes").Find(&user.Notes)

		// Save user to request context and render template
		c.Set("user", user)
		c.HTML(200, template, vars.Default(c))
	}
}

// EditNote returns the page to edit a given note
func EditNote(c *gin.Context) {
	// Acquire variables from request context.
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

	// Load Notes into the user
	db.Model(&user).Association("Notes").Find(&user.Notes)

	// Query db for note
	var note models.Note
	db.Preload("Author").
		Preload("Tags").
		Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		panic(errors.New("Note not found"))
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		panic(errors.New("You are not the owner of this note"))
	}

	// Format note in JSON
	noteJSON, err := json.Marshal(note)
	if err != nil {
		panic("Could not convert note to json")
	}

	// Set request context and render html
	c.Set("user", user)
	c.Set("note", note)
	c.Set("noteJSON", string(noteJSON))
	c.HTML(200, "edit-note.html", vars.Default(c))
}
