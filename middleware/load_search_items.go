package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/search"
)

// LoadSearchItems loads items that match a given query into the request context
func LoadSearchItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get query from request
		q := c.DefaultQuery("q", "")

		// Search db
		response, err := search.Search(q)
		if err != nil {
			c.AbortWithError(500, err)
		}

		// Set searchItems to request context
		c.Set("searchItems", response)
		c.Set("query", q)

		c.Next()
	}
}
