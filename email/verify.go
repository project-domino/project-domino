package email

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
	"github.com/spf13/viper"
)

// SendVerification creates a verification code and
// sends a verification email to a user
func SendVerification(db *gorm.DB, u models.User) error {
	if u.ID == 0 {
		return errors.InternalError
	}

	// Generate random code
	code := make([]byte, 32)
	if _, err := rand.Read(code); err != nil {
		return err
	}
	codeString := base64.RawURLEncoding.EncodeToString(code)

	var err error

	// Set expired field of current verification codes
	err = db.
		Table("email_verification_codes").
		Where("user_id = ?", u.ID).
		Updates(map[string]interface{}{"expired": true}).Error

	if err != nil {
		return err
	}

	err = db.
		Create(&models.EmailVerificationCode{
			User:    u,
			Email:   u.Email,
			Code:    codeString,
			Expired: false,
		}).Error

	if err != nil {
		return err
	}

	verificationEmail := newVerificationEmail(u, codeString)
	if err := Send(verificationEmail, db); err != nil {
		return err
	}

	return nil
}

func newVerificationEmail(user models.User, code string) models.Email {
	bodyFormat := "Please click the link below to verify your email address:\n" +
		"<a href=\"%s\">Verify</a>"
	verifyLink := viper.GetString("http.hostname") + "/email/verify/" + code

	return models.Email{
		User:    user,
		Subject: "Verify your email address",
		Body:    fmt.Sprintf(bodyFormat, verifyLink),
	}
}
