package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// SearchTags searches for tags that match a certain search query
func SearchTags(c *gin.Context) {
	// Acquire db handle from request context.
	db := c.MustGet("db").(*gorm.DB)

	// Acquire variables from request
	query := c.DefaultQuery("q", "")

	var tags []models.Tag
	if query != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", query)

		// Query db
		db.Limit(10).
			Where("name LIKE ?", sqlString).
			Or("description LIKE ?", sqlString).Find(&tags)
	}

	// Return tags in JSON
	c.JSON(http.StatusOK, tags)
}
