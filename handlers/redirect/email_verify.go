package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/handlers"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/notifications"
)

// EmailVerify handles a user's request to verify their email
// Redirects to the home page if successful. Otherwise back to verification page.
func EmailVerify(c *gin.Context) {
	// Get vars from request context
	user := c.MustGet("user").(models.User)
	code := c.Param("verificationCode")

	if user.EmailVerified == true {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// Find verification code in db
	var verification models.EmailVerificationCode
	if err := db.DB.
		Where(&models.EmailVerificationCode{
			UserID:  user.ID,
			Email:   user.Email,
			Code:    code,
			Expired: false,
		}).
		First(&verification).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			handlers.Render(c, "email-verify-error.html", "")
		} else {
			errors.DB.Apply(c)
		}
		return
	}

	tx := db.DB.Begin()
	var err error

	user.EmailVerified = true
	verification.Expired = true

	err = tx.Save(&user).Error
	err = tx.Save(&verification).Error
	err = notifications.RemoveEmailVerifyNotifications(tx, user)

	if err != nil {
		tx.Rollback()
		errors.DB.Apply(c)
		return
	}

	tx.Commit()
	c.Redirect(http.StatusFound, "/")
}
