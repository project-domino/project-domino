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

	// Verify request validity
	if err, ok := verifyNoteRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Create and save note
	newNote := models.Note{
		Title:     requestVars.Title,
		Body:      requestVars.Body,
		Author:    user,
		Published: false,
		Tags:      db.GetTags(requestVars.Tags),
	}
	db.DB.Create(&newNote)

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

	// Clear current note-tag relationships
	db.DB.Model(&note).Association("Tags").Clear()

	// Save note
	note.Title = requestVars.Title
	note.Body = requestVars.Body
	note.Tags = db.GetTags(requestVars.Tags)
	note.Published = requestVars.Publish

	db.DB.Save(&note)

	// Return note in JSON
	c.JSON(http.StatusOK, note)
}

// verifyNoteRequest verifies if values in note request are okay
// returns error and ok value
func verifyNoteRequest(request NoteRequest) (*errors.Error, bool) {
	// check for missing parameters
	if request.Body == "" || request.Title == "" {
		return errors.MissingParameters, false
	}

	return &errors.Error{}, true
}
