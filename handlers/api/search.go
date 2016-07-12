package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// SearchTags searches for tags that match a certain search query
func SearchTags(c *gin.Context) {
	// Return tags in JSON
	tags, err := tagSearch(c.DefaultQuery("q", ""))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, tags)
}

// SearchNotes searches for notes that match a certain search query
func SearchNotes(c *gin.Context) {
	// Return notes in JSON
	notes, err := noteSearch(c.DefaultQuery("q", ""))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, notes)
}

// TODO refine search functions
func noteSearch(q string) ([]models.Note, error) {
	var notes []models.Note
	if q != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", q)

		// Query db
		if err := db.DB.Limit(10).
			Where("title LIKE ?", sqlString).
			Find(&notes).Error; err != nil {
			return notes, err
		}
	}
	return notes, nil
}

func tagSearch(q string) ([]models.Tag, error) {
	var tags []models.Tag
	if q != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", q)

		// Query db
		if err := db.DB.Limit(10).
			Where("name LIKE ?", sqlString).
			Or("description LIKE ?", sqlString).
			Find(&tags).Error; err != nil {
			return tags, err
		}
	}
	return tags, nil
}
