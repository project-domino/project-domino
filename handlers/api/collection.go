package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
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
		Notes:       db.GetNotes(requestVars.Notes),
		Author:      user,
		Published:   false,
		Tags:        db.GetTags(requestVars.Tags),
	}
	db.DB.Create(&newCollection)

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
	db.DB.Preload("Author").Where("id = ?", collectionID).First(&collection)
	if collection.ID == 0 {
		errors.NoteNotFound.Apply(c)
		return
	}

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		errors.NotCollectionOwner.Apply(c)
		return
	}

	// Clear current collection-tag and collection-note relationships
	db.DB.Model(&collection).Association("Tags").Clear()
	db.DB.Model(&collection).Association("Notes").Clear()

	// Edit and save collection
	collection.Title = requestVars.Title
	collection.Description = requestVars.Description
	collection.Notes = db.GetNotes(requestVars.Notes)
	collection.Author = user
	collection.Published = requestVars.Publish
	collection.Tags = db.GetTags(requestVars.Tags)

	db.DB.Save(&collection)

	// Return collection in JSON
	c.JSON(http.StatusOK, collection)
}
