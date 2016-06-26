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
	c.JSON(http.StatusOK, tagSearch(c.DefaultQuery("q", "")))
}

// SearchNotes searches for notes that match a certain search query
func SearchNotes(c *gin.Context) {
	// Return notes in JSON
	c.JSON(http.StatusOK, noteSearch(c.DefaultQuery("q", "")))
}

// TODO refine search functions
func noteSearch(q string) []models.Note {
	var notes []models.Note
	if q != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", q)

		// Query db
		db.DB.Limit(10).
			Where("title LIKE ?", sqlString).Find(&notes)
	}
	return notes
}

func tagSearch(q string) []models.Tag {
	var tags []models.Tag
	if q != "" {
		// Create SQL search string
		sqlString := fmt.Sprintf("%%%s%%", q)

		// Query db
		db.DB.Limit(10).
			Where("name LIKE ?", sqlString).
			Or("description LIKE ?", sqlString).Find(&tags)
	}
	return tags
}
