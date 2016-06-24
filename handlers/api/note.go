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

	sanitizedRequest, err := sanitizeRequest(requestVars)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Query db for tags
	var tags []models.Tag
	util.DB.Where("id in (?)", sanitizedRequest.Tags).Find(&tags)

	// Create and save note
	newNote := models.Note{
		Title:     sanitizedRequest.Title,
		Body:      sanitizedRequest.Body,
		Author:    user,
		Published: false,
		Tags:      tags,
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

	sanitizedRequest, err := sanitizeRequest(requestVars)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Query db for tags
	var tags []models.Tag
	util.DB.Where("id in (?)", sanitizedRequest.Tags).Find(&tags)

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
	note.Title = sanitizedRequest.Title
	note.Body = sanitizedRequest.Body
	note.Tags = tags
	note.Published = sanitizedRequest.Publish

	util.DB.Save(&note)

	// Return note in JSON
	c.JSON(http.StatusOK, note)
}

func sanitizeRequest(request NoteRequest) (NoteRequest, error) {
	// TODO 10 seems like a good number for max tags?
	// Especially if some tags depend on others.
	if len(request.Tags) > 10 {
		return NoteRequest{}, errors.New("Too many tags.")
	}

	// Remove duplicate tags
	var tempTags []uint
	for _, tag := range request.Tags {
		if !contains(tempTags, tag) {
			tempTags = append(tempTags, tag)
		}
	}
	request.Tags = tempTags

	return request, nil
}

func contains(set []uint, element uint) bool {
	for _, e := range set {
		if e == element {
			return true
		}
	}
	return false
}
