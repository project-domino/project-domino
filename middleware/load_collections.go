package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// These constants are the valid values by in LoadCollections
const (
	LoadCollectionsAuthor      string = "author"
	LoadCollectionsRequestUser        = "requestUser"
	LoadCollectionsUserUpvote         = "requestUserUpvote"
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

		preloadedDB = preloadedDB.
			Limit(items).
			Offset((page - 1) * items).
			Order("updated_at desc")

		var collections []models.Collection
		var err error

		switch by {
		case LoadCollectionsAuthor:
			user := c.MustGet("pageUser").(models.User)
			err = preloadedDB.
				Where("author_id = ?", user.ID).
				Where("published = ?", true).
				Find(&collections).
				Error
		case LoadCollectionsUserUpvote:
			user := c.MustGet("pageUser").(models.User)
			err = preloadedDB.
				Model(&user).
				Association("UpvoteCollections").
				Find(&collections).
				Error
		case LoadCollectionsRequestUser:
			user := c.MustGet("user").(models.User)
			err = preloadedDB.
				Where("author_id = ?", user.ID).
				Find(&collections).
				Error
		default:
			errors.InternalError.Apply(c)
			return
		}

		if err != nil {
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
