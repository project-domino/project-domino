package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// TextbookRequest holds the request object for NewTextbook and EditTextbook
type TextbookRequest struct {
	Collections []uint `json:"collections" binding:"required"`
	Publish     bool   `json:"publish"`
	Tags        []uint `json:"tags" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// NewTextbook creates a collection with specified values
func NewTextbook(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Get request variables
	var requestVars TextbookRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Create and save collection
	newTextbook := models.Textbook{
		Title:       requestVars.Title,
		Description: requestVars.Description,
		Collections: util.GetCollections(requestVars.Collections),
		Author:      user,
		Published:   false,
		Tags:        util.GetTags(requestVars.Tags),
	}
	util.DB.Create(&newTextbook)

	// Return collection in JSON
	c.JSON(http.StatusOK, newTextbook)
}

// EditTextbook edits a textbook with specified values
func EditTextbook(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	textbookID := c.Param("textbookID")

	// Get request variables
	var requestVars TextbookRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Query db for textbook
	var textbook models.Textbook
	util.DB.Preload("Author").Where("id = ?", textbookID).First(&textbook)
	if textbook.ID == 0 {
		c.AbortWithError(404, errors.New("Collection not found"))
		return
	}

	// Check if request user is the owner of the textbook
	if textbook.Author.ID != user.ID {
		c.AbortWithError(403, errors.New("You are not the owner of this textbook"))
		return
	}

	// Clear current textbook-collection and textbook-tag relationships
	util.DB.Model(&textbook).Association("Tags").Clear()
	util.DB.Model(&textbook).Association("Collections").Clear()

	// Edit and save textbook
	textbook.Title = requestVars.Title
	textbook.Description = requestVars.Description
	textbook.Collections = util.GetCollections(requestVars.Collections)
	textbook.Author = user
	textbook.Published = requestVars.Publish
	textbook.Tags = util.GetTags(requestVars.Tags)

	util.DB.Save(&textbook)

	// Return collection in JSON
	c.JSON(http.StatusOK, textbook)
}
