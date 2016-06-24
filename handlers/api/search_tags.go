package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// SearchTags searches for tags that match a certain search query
func SearchTags(c *gin.Context) {
	// Acquire variables from request
	query := c.DefaultQuery("q", "")

	var tags []models.Tag
	if query != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", query)

		// Query db
		util.DB.Limit(10).
			Where("name LIKE ?", sqlString).
			Or("description LIKE ?", sqlString).Find(&tags)
	}

	// Return tags in JSON
	c.JSON(http.StatusOK, tags)
}
