package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// These constants are the valid values by in LoadCollections
const (
	LoadCollectionsAuthor string = "author"
)

// LoadCollections loads collections into the request context with specified objects
// It also loads collectionsJSON, the collections serialized into JSON
// The first argument must be what the collections are loaded by
// "by" can be one of the following options: author
func LoadCollections(by string, objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get list vars
		items, page, varErr := getListVars(c)
		c.Set("items", items)
		if varErr != nil {
			varErr.Apply(c)
			return
		}

		// Set objects to be preloaded to db
		preloadedDB := db.DB
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		switch by {
		case "author":
			// Acquire username from URL
			username := c.Param("username")
			// Find user in db
			var user models.User
			if err := db.DB.
				Where("user_name = ?", username).
				First(&user).
				Error; err != nil {

				if err == gorm.ErrRecordNotFound {
					errors.UserNotFound.Apply(c)
				} else {
					errors.DB.Apply(c)
				}
				return
			}
			// add where statement to query
			preloadedDB = preloadedDB.Where("author_id = ?", user.ID)

		default:
			errors.InternalError.Apply(c)
			return
		}

		// Query for collections
		var collections []models.Collection
		if err := preloadedDB.
			Limit(items).
			Offset((page - 1) * items).
			Find(&collections).
			Error; err != nil {

			errors.DB.Apply(c)
			return
		}

		// Format collections in JSON
		collectionsJSON, err := json.Marshal(collections)
		if err != nil {
			errors.JSON.Apply(c)
			return
		}

		// Set pages
		prevPage, nextPage := getPages(page, len(collections) < items)
		c.Set("currentPage", page)
		c.Set("nextPage", nextPage)
		c.Set("prevPage", prevPage)

		// Add collections and collectionsJSON to request context
		c.Set("collections", collections)
		c.Set("collectionsJSON", string(collectionsJSON))

		c.Next()
	}
}
