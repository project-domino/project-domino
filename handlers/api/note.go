package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// NoteRequest holds the request object for NewNote and EditNote
type NoteRequest struct {
	Body    string `json:"body" binding:"required"`
	Publish bool   `json:"publish"`
	Tags    []uint `json:"tags" binding:"required"`
	Title   string `json:"title" binding:"required"`
}

// NewNote creates a note with a specified values
func NewNote(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Get request variables
	var requestVars NoteRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Create and save note
	newNote := models.Note{
		Title:     requestVars.Title,
		Body:      requestVars.Body,
		Author:    user,
		Published: false,
		Tags:      util.GetTags(requestVars.Tags),
	}
	util.DB.Create(&newNote)

	// Return note in JSON
	c.JSON(http.StatusOK, newNote)
}

// EditNote edits a note with specified values
func EditNote(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

	// Get request variables
	var requestVars NoteRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Query db for note
	var note models.Note
	util.DB.Preload("Author").Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		c.AbortWithError(404, errors.New("Note not found"))
		return
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		c.AbortWithError(403, errors.New("You are not the owner of this note"))
		return
	}

	// Clear current note-tag relationships
	util.DB.Model(&note).Association("Tags").Clear()

	// Save note
	note.Title = requestVars.Title
	note.Body = requestVars.Body
	note.Tags = util.GetTags(requestVars.Tags)
	note.Published = requestVars.Publish

	util.DB.Save(&note)

	// Return note in JSON
	c.JSON(http.StatusOK, note)
}
