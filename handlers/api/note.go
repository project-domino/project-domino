package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	// Acquire variables from request context.
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(models.User)

	// Get request variables
	var requestVars NoteRequest
	if err := c.BindJSON(&requestVars); err != nil {
		panic(err)
	}

	sanitizedRequest, err := sanitizeRequest(requestVars)
	if err != nil {
		panic(err)
	}

	// Query db for tags
	var tags []models.Tag
	db.Where("id in (?)", sanitizedRequest.Tags).Find(&tags)

	// Create and save note
	newNote := models.Note{
		Title:     sanitizedRequest.Title,
		Body:      sanitizedRequest.Body,
		Author:    user,
		Published: false,
		Tags:      tags,
	}
	db.Create(&newNote)

	// Return note in JSON
	c.JSON(http.StatusOK, newNote)
}

// EditNote edits a note with specified values
func EditNote(c *gin.Context) {
	// Acquire variables from request context.
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")

	// Get request variables
	var requestVars NoteRequest
	if err := c.BindJSON(&requestVars); err != nil {
		panic(err)
	}

	sanitizedRequest, err := sanitizeRequest(requestVars)
	if err != nil {
		panic(err)
	}

	// Query db for tags
	var tags []models.Tag
	db.Where("id in (?)", sanitizedRequest.Tags).Find(&tags)

	// Query db for note
	var note models.Note
	db.Preload("Author").Where("id = ?", noteID).First(&note)
	if note.ID == 0 {
		panic(errors.New("Note not found"))
	}

	// Check if request user is the owner of the note
	if note.Author.ID != user.ID {
		panic(errors.New("You are not the owner of this note"))
	}

	// Clear current note-tag relationships
	db.Model(&note).Association("Tags").Clear()

	// Save note
	note.Title = sanitizedRequest.Title
	note.Body = sanitizedRequest.Body
	note.Tags = tags
	note.Published = sanitizedRequest.Publish

	db.Save(&note)

	// Return note in JSON
	c.JSON(http.StatusOK, note)
}

func sanitizeRequest(request NoteRequest) (NoteRequest, error) {
	// TODO 10 seems like a good number for max tags?
	// Especially if some tags depend on others.
	if len(request.Tags) > 10 {
		return NoteRequest{}, errors.New("To many tags.")
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
