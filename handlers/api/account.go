package api

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// EditName handles a users request to change their name
func EditName(c *gin.Context) {
	// Get values from request context
	user := c.MustGet("user").(models.User)
	name := c.PostForm("name")

	// Set value and save to db
	user.Name = name

	if err := db.DB.Save(&user).Error; err != nil {
		errors.DB.Apply(c)
		return
	}
}

// EditEmail handles a users request to change their email
func EditEmail(c *gin.Context) {
	// Get values from request context
	// user := c.MustGet("user")
	// email := c.PostForm("email")
}
