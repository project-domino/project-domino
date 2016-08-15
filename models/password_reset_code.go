package models

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/jinzhu/gorm"
)

// An PasswordResetCode holds the verification code for a user to reset their
// password
type PasswordResetCode struct {
	gorm.Model

	UserID uint
	User   User

	Code string

	Expired bool
}

// NewPasswordResetCode returns a new password reset code
func NewPasswordResetCode(user User) (PasswordResetCode, error) {
	var resetCode PasswordResetCode

	resetCode.User = user
	resetCode.Expired = false

	// Generate random code
	code := make([]byte, 16)
	if _, err := rand.Read(code); err != nil {
		return resetCode, err
	}
	codeString := base64.RawURLEncoding.EncodeToString(code)

	resetCode.Code = codeString

	return resetCode, nil
}
