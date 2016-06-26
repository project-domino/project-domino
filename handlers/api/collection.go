package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// CollectionRequest holds the request object for NewCollection and EditCollection
type CollectionRequest struct {
	Notes       []uint `json:"notes" binding:"required"`
	Publish     bool   `json:"publish"`
	Tags        []uint `json:"tags" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// NewCollection creates a collection with specified values
func NewCollection(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Get request variables
	var requestVars CollectionRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Create and save collection
	newCollection := models.Collection{
		Title:       requestVars.Title,
		Description: requestVars.Description,
		Notes:       util.GetNotes(requestVars.Notes),
		Author:      user,
		Published:   false,
		Tags:        util.GetTags(requestVars.Tags),
	}
	util.DB.Create(&newCollection)

	// Return collection in JSON
	c.JSON(http.StatusOK, newCollection)
}

// EditCollection edits a collection with specified values
func EditCollection(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	collectionID := c.Param("collectionID")

	// Get request variables
	var requestVars CollectionRequest
	if err := c.BindJSON(&requestVars); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Query db for collection
	var collection models.Collection
	util.DB.Preload("Author").Where("id = ?", collectionID).First(&collection)
	if collection.ID == 0 {
		c.AbortWithError(404, errors.New("Note not found"))
		return
	}

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		c.AbortWithError(403, errors.New("You are not the owner of this collection"))
		return
	}

	// Clear current collection-tag and collection-note relationships
	util.DB.Model(&collection).Association("Tags").Clear()
	util.DB.Model(&collection).Association("Notes").Clear()

	// Edit and save collection
	collection.Title = requestVars.Title
	collection.Description = requestVars.Description
	collection.Notes = util.GetNotes(requestVars.Notes)
	collection.Author = user
	collection.Published = requestVars.Publish
	collection.Tags = util.GetTags(requestVars.Tags)

	util.DB.Save(&collection)

	// Return collection in JSON
	c.JSON(http.StatusOK, collection)
}