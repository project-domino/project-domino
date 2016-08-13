package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadComment loads a comment into the request context with specified objects
// :commentID must be in the URL
func LoadComment(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire commentID from URL
		commentID := c.Param("commentID")

		// Set objects to be preloaded to db
		preloadedDB := db.DB.Where("id = ?", commentID)
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// Query for comment
		var comment models.Comment
		if err := preloadedDB.First(&comment).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors.CommentNotFound.Apply(c)
				return
			}
			errors.DB.Apply(c)
			return
		}

		// Add comment to request context
		c.Set("comment", comment)

		c.Next()
	}
}
