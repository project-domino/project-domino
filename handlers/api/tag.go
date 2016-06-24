package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/util"
)

// NewTag creates a tag with a specified values
func NewTag(c *gin.Context) {
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
	util.DB.Where("name = ?", name).Find(&checkTags)

	// If tag exists, return error
	if len(checkTags) != 0 {
		c.AbortWithError(400, errors.New("Tag with same name already exists"))
		return
	}

	// Get request user
	user := c.MustGet("user").(models.User)

	// Create and save tag
	util.DB.Create(&models.Tag{
		Name:        name,
		Description: description,
		Author:      user,
	})
}
