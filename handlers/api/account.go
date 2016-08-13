package api

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/email"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/notifications"
)

// EditUser handles a users request to edit their info
func EditUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	name := c.PostForm("name")
	email := c.PostForm("email")

	var err error

	if name != user.Name {
		err = editName(&user, name)
	}

	if email != user.Email {
		err = editEmail(&user, email)
	}

	if err != nil {
		errors.DB.Apply(c)
		return
	}
}

// SendEmailVerification sends an email verification to a user
func SendEmailVerification(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	if err := email.SendVerification(db.DB, user); err != nil {
		errors.DB.Apply(c)
	}
}

func editName(user *models.User, name string) error {
	user.Name = name

	err := db.DB.Save(user).Error

	return err
}

func editEmail(user *models.User, e string) error {
	user.Email = e
	user.EmailVerified = false

	tx := db.DB.Begin()

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := email.SendVerification(tx, *user); err != nil {
		tx.Rollback()
		return err
	}

	if err := notifications.EmailVerify(tx, *user); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
