package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadCollection loads a collection into the request context with specified objects
// It also loads collectionJSON, the collection object serialized into JSON
// :collectionID must be in the URL
func LoadCollection(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire collectionID from URL
		collectionID := c.Param("collectionID")

		// Set objects to be preloaded to db
		preloadedDB := db.DB.Where("id = ?", collectionID)
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// Query for collection
		var collection models.Collection
		if err := preloadedDB.First(&collection).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors.CollectionNotFound.Apply(c)
				return
			}
			c.AbortWithError(500, err)
			return
		}

		// Load notes into the collection
		if err := db.LoadCollectionNotes(&collection); err != nil {
			c.AbortWithError(500, err)
			return
		}

		// Format collection in JSON
		collectionJSON, err := json.Marshal(collection)
		if err != nil {
			c.Error(err)
			errors.JSON.Apply(c)
			return
		}

		// Add collection and collectionJSON to request context
		c.Set("collection", collection)
		c.Set("collectionJSON", string(collectionJSON))

		c.Next()
	}
}
