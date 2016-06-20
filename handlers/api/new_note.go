package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// NewNoteRequest holds the request object for NewNote
type NewNoteRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Tags  []uint `json:"tags" binding:"required"`
}

// NewNote creates a note with a specified values
func NewNote(c *gin.Context) {
	// Acquire db handle from request context.
	db := c.MustGet("db").(*gorm.DB)

	// Get request variables
	var requestVars NewNoteRequest
	if c.BindJSON(&requestVars) != nil {
		panic(errors.New("Invalid Parameters"))
	}

	// TODO 10 seems like a good number for max tags?
	// Especially if some tags depend on others.
	if len(requestVars.Tags) > 10 {
		panic(errors.New("Invalid Parameters"))
	}

	// Remove duplicate tags
	var tempTags []uint
	for _, tag := range requestVars.Tags {
		if !contains(tempTags, tag) {
			tempTags = append(tempTags, tag)
		}
	}
	// Query db for tags
	var tags []models.Tag
	db.Where("id in (?)", tempTags).Find(&tags)

	// Get request user
	user := c.MustGet("user").(models.User)

	// Create and save note
	db.Create(&models.Note{
		Title:     requestVars.Title,
		Body:      requestVars.Body,
		Author:    user,
		Published: false,
		Tags:      tags,
	})
}

func contains(set []uint, element uint) bool {
	for _, e := range set {
		if e == element {
			return true
		}
	}
	return false
}