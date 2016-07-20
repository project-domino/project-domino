package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
)

// Search handles searching the db
func Search(c *gin.Context) {
	// Get search results from context
	response, ok := c.Get("searchResult")
	if !ok {
		errors.InternalError.Apply(c)
		return
	}

	// Send back response in json
	c.JSON(http.StatusOK, response)
}
