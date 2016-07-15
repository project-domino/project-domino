package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// NoteRequest holds the request object for NewNote and EditNote
type NoteRequest struct {
	Body        string `json:"body" binding:"required"`
	Description string `json:"description" binding:"required"`
	Publish     bool   `json:"publish"`
	Tags        []uint `json:"tags" binding:"required"`
	Title       string `json:"title" binding:"required"`
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

	// Verify request validity
	if err, ok := verifyNoteRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Get request tags
	tags, err := db.GetTags(requestVars.Tags)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Create and save note
	newNote := models.Note{
		Title:       requestVars.Title,
		Body:        requestVars.Body,
		Description: requestVars.Description,
		Author:      user,
		Published:   requestVars.Publish,
		Tags:        tags,
	}
	if err := db.DB.Create(&newNote).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}

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

	// Verify request validity
	if err, ok := verifyNoteRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Get request tags
	tags, err := db.GetTags(requestVars.Tags)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Query db for note
	var note models.Note
	db.DB.Preload("Author").Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		errors.NoteNotFound.Apply(c)
		return
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		errors.NotNoteOwner.Apply(c)
		return
	}

	// Create transaction to save note
	tx := db.DB.Begin()

	// Save note-tag relationships
	if err := tx.Model(&note).
		Association("Tags").
		Replace(tags).Error; err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	// Save note
	note.Title = requestVars.Title
	note.Body = requestVars.Body
	note.Description = requestVars.Description
	note.Published = requestVars.Publish

	if err := tx.Save(&note).Error; err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	tx.Commit()

	// Return note in JSON
	c.JSON(http.StatusOK, note)
}

// verifyNoteRequest verifies if values in note request are okay
// returns error and ok value
func verifyNoteRequest(request NoteRequest) (*errors.Error, bool) {
	// check for missing parameters
	if request.Body == "" || request.Title == "" || request.Description == "" {
		return errors.MissingParameters, false
	}

	// Check description size
	if len(request.Description) > 250 {
		return errors.BadParameters, false
	}

	return &errors.Error{}, true
}
