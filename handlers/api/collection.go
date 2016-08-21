package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		errors.BadParameters.Apply(c)
		return
	}

	// Verify request validity
	if err, ok := verifyCollectionRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Get request tags
	tags, err := db.GetTags(requestVars.Tags)
	if err != nil {
		errors.DB.Apply(c)
		return
	}

	// Create transaction to create collection
	tx := db.DB.Begin()

	// Create and save collection
	newCollection := models.Collection{
		Title:       requestVars.Title,
		Description: requestVars.Description,
		Author:      user,
		Published:   requestVars.Publish,
		Tags:        tags,
	}
	if err := tx.Create(&newCollection).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	// Save collection-note relationships
	for i, noteID := range requestVars.Notes {
		var relation models.CollectionNote
		relation.Collection = newCollection
		relation.NoteID = noteID
		relation.Order = uint(i) + 1

		if err := tx.Create(&relation).Error; err != nil {
			tx.Rollback()
			errors.DB.Apply(c)
			return
		}
	}

	tx.Commit()

	// Update searchtext field
	if err := db.UpdateCollectionSearchText(newCollection.ID); err != nil {
		errors.DB.Apply(c)
		return
	}

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
		errors.BadParameters.Apply(c)
		return
	}

	// Verify request validity
	if err, ok := verifyCollectionRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Get request tags
	tags, err := db.GetTags(requestVars.Tags)
	if err != nil {
		errors.DB.Apply(c)
		return
	}

	// Query db for collection
	var collection models.Collection
	if err := db.DB.Preload("Author").
		Where("id = ?", collectionID).First(&collection).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.CollectionNotFound.Apply(c)
			return
		}
		errors.DB.Apply(c)
		return
	}

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		errors.NotCollectionOwner.Apply(c)
		return
	}

	// Create transaction to save collection
	tx := db.DB.Begin()

	// Clear current collection-note relationships
	if err := tx.Where("collection_id = ?", collection.ID).
		Delete(models.CollectionNote{}).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	// Save collection-tag relationships
	if err := tx.Model(&collection).
		Association("Tags").
		Replace(tags).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	// Edit and save collection
	collection.Title = requestVars.Title
	collection.Description = requestVars.Description
	collection.Author = user
	collection.Published = requestVars.Publish

	if err := tx.Save(&collection).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	// Save collection-note relationships
	for i, noteID := range requestVars.Notes {
		var relation models.CollectionNote
		relation.Collection = collection
		relation.NoteID = noteID
		relation.Order = uint(i) + 1

		if err := tx.Create(&relation).Error; err != nil {
			tx.Rollback()
			errors.DB.Apply(c)
			return
		}
	}

	tx.Commit()

	// Update searchtext field
	if err := db.UpdateCollectionSearchText(collection.ID); err != nil {
		errors.DB.Apply(c)
		return
	}

	// Return collection in JSON
	c.JSON(http.StatusOK, collection)
}

// verifyCollectionRequest verifies if values in collection request are okay
// returns error and ok value
func verifyCollectionRequest(request CollectionRequest) (*errors.Error, bool) {
	// Check for missing parameters
	if request.Description == "" || request.Title == "" || len(request.Notes) == 0 {
		return errors.MissingParameters, false
	}

	// Check description size
	if len(request.Description) > 500 {
		return errors.BadParameters, false
	}

	// Check for duplicate notes
	var temp []uint
	for _, e := range request.Notes {
		if !contains(temp, e) {
			temp = append(temp, e)
		} else {
			return errors.BadParameters, false
		}
	}

	// Verify notes exist
	if err := db.VerifyNotes(request.Notes); err != nil {
		return errors.NoteNotFound, false
	}

	return &errors.Error{}, true
}

// contains checks if array "a" contains uint "e"
func contains(a []uint, e uint) bool {
	for _, i := range a {
		if i == e {
			return true
		}
	}
	return false
}
