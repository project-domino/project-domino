package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// NewTag creates a tag with a specified values
func NewTag(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	name := c.PostForm("name")
	description := c.PostForm("description")

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

	tag := models.Tag{
		Name:        name,
		Description: description,
		Author:      user,
	}

	if err := db.DB.
		Create(&tag).
		Error; err != nil {
		errors.DB.Apply(c)
		return
	}

	c.JSON(http.StatusOK, tag)
}
