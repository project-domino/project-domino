package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/search"
)

// Search handles searching the db
func Search(c *gin.Context) {
	// Get query from request
	q := c.DefaultQuery("q", "")

	// Search db
	response, err := search.Search(q)
	if err != nil {
		c.AbortWithError(500, err)
	}

	// Send back response in json
	c.JSON(http.StatusOK, response)
}
