package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/email"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// ResetPassword handles a user's request to reset their password
func ResetPassword(c *gin.Context) {
	userName := c.PostForm("userName")
	resetCodeString := c.PostForm("resetCode")
	password := c.PostForm("password")
	retypePassword := c.PostForm("retypePassword")

	if userName == "" || resetCodeString == "" || password == "" || retypePassword == "" {
		errors.BadParameters.Apply(c)
		return
	}

	if password != retypePassword {
		errors.PasswordsDoNotMatch.Apply(c)
		return
	}

	var user models.User
	if err := db.DB.
		Where("user_name = ?", userName).
		First(&user).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			errors.UserNotFound.Apply(c)
		} else {
			errors.DB.Apply(c)
		}

		return
	}

	var resetCode models.PasswordResetCode
	if err := db.DB.
		Where("user_id = ?", user.ID).
		Where("code = ?", resetCodeString).
		Where("expired = ?", false).
		First(&resetCode).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			errors.ResetCodeNotFound.Apply(c)
		} else {
			errors.DB.Apply(c)
		}

		return
	}

	resetCode.Expired = true

	if err := user.SetPassword(password); err != nil {
		c.AbortWithError(500, err)
		return
	}

	tx := db.DB.Begin()

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}
	if err := tx.Save(&resetCode).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	tx.Commit()
}

// SendPasswordResetCode sends a password reset code to a given user's email
func SendPasswordResetCode(c *gin.Context) {
	userName := c.PostForm("userName")

	var user models.User

	if err := db.DB.
		Where("user_name = ?", userName).
		First(&user).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			errors.UserNotFound.Apply(c)
		} else {
			errors.DB.Apply(c)
		}

		return
	}

	if !user.EmailVerified {
		errors.EmailNotVerified.Apply(c)
		return
	}

	passwordResetCode, err := models.NewPasswordResetCode(user)
	if err != nil {
		errors.InternalError.Apply(c)
		return
	}

	tx := db.DB.Begin()

	if err := tx.
		Table("password_reset_codes").
		Where("user_id = ?", user.ID).
		Updates(map[string]interface{}{"expired": true}).
		Error; err != nil {

		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	if err := tx.Create(&passwordResetCode).Error; err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	codeEmail := models.Email{
		User:    user,
		Subject: "Password Reset",
		Body:    fmt.Sprintf("Your password reset code is: %s", passwordResetCode.Code),
	}
	if err := email.Send(codeEmail, tx); err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	tx.Commit()
}
