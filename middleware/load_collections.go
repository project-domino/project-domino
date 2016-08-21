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
	LoadCollectionsAuthor      string = "author"
	LoadCollectionsRequestUser        = "requestUser"
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
		case LoadCollectionsAuthor:
			var user models.User
			if err := db.DB.Where("user_name = ?", c.Param("username")).First(&user).
				Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					errors.UserNotFound.Apply(c)
				} else {
					errors.DB.Apply(c)
				}
				return
			}
			preloadedDB = preloadedDB.
				Where("author_id = ?", user.ID).
				Where("published = ?", true)
		case LoadCollectionsRequestUser:
			user := c.MustGet("user").(models.User)
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
			Order("updated_at desc").
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
