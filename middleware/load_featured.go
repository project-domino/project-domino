package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// LoadFeaturedItems loads all featured notes and collections into the
// request context
func LoadFeaturedItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var notes []models.Note
		var collections []models.Collection

		preloadedDB := db.DB.
			Where("featured = ?", true).
			Order("updated_at desc").
			Preload("Tags")

		preloadedDB.
			Find(&notes)

		preloadedDB.
			Find(&collections)

		c.Set("featuredNotes", notes)
		c.Set("featuredCollections", collections)

		c.Next()
	}
}
