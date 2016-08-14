package api

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// ChangePassword handles a user's request to reset their password
func ChangePassword(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	oldPass := c.PostForm("oldPassword")
	newPass := c.PostForm("newPassword")
	newRetypePass := c.PostForm("newRetypePassword")

	if oldPass == "" || newPass == "" || newRetypePass == "" {
		errors.MissingParameters.Apply(c)
		return
	}

	if newPass != newRetypePass {
		errors.BadParameters.Apply(c)
		return
	}

	oldPassValid := user.CheckPassword(oldPass)
	if !oldPassValid {
		errors.InvalidCredentials.Apply(c)
		return
	}

	if err := user.SetPassword(newPass); err != nil {
		c.AbortWithError(500, err)
		return
	}

	if err := db.DB.Save(&user).Error; err != nil {
		errors.DB.Apply(c)
		return
	}
}
