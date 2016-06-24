package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// NewTag creates a tag with a specified values
func NewTag(c *gin.Context) {
	// Acquire db handle from request context.
	db := c.MustGet("db").(*gorm.DB)

	// Get request variables
	name := c.PostForm("name")
	description := c.PostForm("description")

	// Check for valid values
	if name == "" || description == "" {
		c.AbortWithError(400, errors.New("There are empty fields"))
		return
	}

	// Check if tag exists
	var checkTags []models.Tag
	db.Where("name = ?", name).Find(&checkTags)

	// If tag exists, return error
	if len(checkTags) != 0 {
		c.AbortWithError(400, errors.New("Tag with same name already exists"))
		return
	}

	// Get request user
	user := c.MustGet("user").(models.User)

	// Create and save tag
	db.Create(&models.Tag{
		Name:        name,
		Description: description,
		Author:      user,
	})
}
