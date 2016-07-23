package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// VerifyCollectionPublic checks if the collection in the request context is public
func VerifyCollectionPublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire collection from request context
		collection := c.MustGet("collection").(models.Collection)

		// Check if the collection is public
		if !collection.Published {
			errors.CollectionNotFound.Apply(c)
			return
		}

		c.Next()
	}
}
