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
		errors.BadParameters.Apply(c)
		return
	}

	// Verify request validity
	if err, ok := verifyCollectionRequest(requestVars); !ok {
		err.Apply(c)
		return
	}

	// Create and save collection
	newCollection := models.Collection{
		Title:       requestVars.Title,
		Description: requestVars.Description,
		Author:      user,
		Published:   false,
		Tags:        db.GetTags(requestVars.Tags),
	}
	db.DB.Create(&newCollection)

	// Save collection-note relationships
	saveNoteRelations(newCollection, requestVars.Notes)

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

	// Query db for collection
	var collection models.Collection
	db.DB.Preload("Author").Where("id = ?", collectionID).First(&collection)
	if collection.ID == 0 {
		errors.CollectionNotFound.Apply(c)
		return
	}

	// Check if request user is the owner of the collection
	if collection.Author.ID != user.ID {
		errors.NotCollectionOwner.Apply(c)
		return
	}

	// Clear current collection-note relationships
	db.DB.Where("collection_id = ?", collection.ID).Delete(models.CollectionNote{})

	// Save collection-tag relationships
	db.DB.Model(&collection).Association("Tags").Replace(db.GetTags(requestVars.Tags))

	// Edit and save collection
	collection.Title = requestVars.Title
	collection.Description = requestVars.Description
	collection.Author = user
	collection.Published = requestVars.Publish

	db.DB.Save(&collection)

	// Save collection-note relationships
	saveNoteRelations(collection, requestVars.Notes)

	// Return collection in JSON
	c.JSON(http.StatusOK, collection)
}

// verifyCollectionRequest verifies if values in collection request are okay
// returns error and ok value
func verifyCollectionRequest(request CollectionRequest) (*errors.Error, bool) {
	// check for missing parameters
	if request.Description == "" || request.Title == "" || len(request.Notes) == 0 {
		return errors.MissingParameters, false
	}

	// check for duplicate notes
	var temp []uint
	for _, e := range request.Notes {
		if !contains(temp, e) {
			temp = append(temp, e)
		} else {
			return errors.BadParameters, false
		}
	}

	// verify notes exist
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

// saveNoteRelations saves the relationships between a collection and notes
func saveNoteRelations(collection models.Collection, notes []uint) {
	for i, noteID := range notes {
		var relation models.CollectionNote
		relation.Collection = collection
		relation.NoteID = noteID
		relation.Order = uint(i) + 1

		db.DB.Create(&relation)
	}
}
