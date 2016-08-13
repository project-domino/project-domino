package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// NewTag creates a tag with a specified values
func NewTag(c *gin.Context) {
	// Get request variables
	name := c.PostForm("name")
	description := c.PostForm("description")

	// Check for valid values
	if name == "" || description == "" {
		errors.MissingParameters.Apply(c)
		return
	}

	// Check if tag exists
	var checkTags []models.Tag
	if err := db.DB.Where("name = ?", name).
		Find(&checkTags).Error; err != nil && err != gorm.ErrRecordNotFound {
		errors.DB.Apply(c)
		return
	}

	// If tag exists, return error
	if len(checkTags) != 0 {
		errors.TagExists.Apply(c)
		return
	}

	// Get request user
	user := c.MustGet("user").(models.User)

	// Create and save tag
	if err := db.DB.Create(&models.Tag{
		Name:        name,
		Description: description,
		Author:      user,
	}).Error; err != nil {
		errors.DB.Apply(c)
		return
	}
}
