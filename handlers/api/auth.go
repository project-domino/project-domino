package api

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// Login handles requests to log a user in.
// If credentials are valid, sets an auth cookie.
func Login(c *gin.Context) {
	// Get needed variables from request.
	userName := c.PostForm("userName")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// If there are blank fields, return bad request
	if password == "" || (userName == "" && email == "") {
		errors.MissingParameters.Apply(c)
		return
	}

	// Find user in the database
	var users []models.User
	if err := db.DB.Limit(1).
		Where("email = ?", email).
		Or("user_name = ?", userName).
		Find(&users).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		errors.DB.Apply(c)
		return
	}

	// If there are no users with specified username/email return error
	if len(users) == 0 {
		errors.InvalidCredentials.Apply(c)
		return
	}

	// Otherwise, check password and assign cookie
	user := users[0]
	if !user.CheckPassword(password) {
		errors.InvalidCredentials.Apply(c)
		return
	}

	// TODO unlegacy
	AuthCookie(c, user)
}

// Logout handles requests to log a user out.
func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "auth",
		Path:    "/",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

// Register handles requests to handle a new user.
func Register(c *gin.Context) {
	// Get needed variables from request.
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	retypePassword := c.PostForm("retypePassword")

	// Check if the request is missing needed parameters
	if userName == "" || password == "" || retypePassword == "" {
		errors.MissingParameters.Apply(c)
		return
	}

	// Check if password matches retype password
	if retypePassword != password {
		errors.PasswordsDoNotMatch.Apply(c)
		return
	}

	// Check if other users have the same email or userName
	var checkUsers []models.User
	if err := db.DB.
		Or("user_name = ?", userName).
		Find(&checkUsers).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		errors.DB.Apply(c)
		return
	}

	if len(checkUsers) != 0 {
		errors.UserExists.Apply(c)
		return
	}

	// Create the user.
	user := models.User{
		UserName: userName,
		Type:     models.General,
	}
	if err := user.SetPassword(password); err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Add user to database.
	if err := db.DB.Create(&user).Error; err != nil {
		errors.DB.Apply(c)
		return
	}

	// Set an auth cookie for the user
	// TODO unlegacy
	AuthCookie(c, user)
	c.Redirect(http.StatusFound, "/")
}

// AuthCookie creates an authentication token and sends it to the client.
func AuthCookie(c *gin.Context, user models.User) {
	var authToken models.AuthToken
	var err error

	if strings.Contains(c.Request.UserAgent(), "Domino") {
		authToken, err = newAuthToken(user)
	} else {
		authToken, err = newAuthTokenExpires(user, time.Now().Add(time.Hour*24*7))
	}

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if err := db.DB.Create(&authToken).Error; err != nil {
		errors.DB.Apply(c)
		return
	}

	// Create cookie
	cookie := http.Cookie{
		Name:    "auth",
		Path:    "/",
		Value:   authToken.Token,
		Expires: authToken.Expires,
	}
	http.SetCookie(c.Writer, &cookie)
}

// newAuthToken creates a new AuthToken that never expires.
func newAuthToken(user models.User) (models.AuthToken, error) {
	// We can't use 1<<63 because reasons.
	return newAuthTokenExpires(user, time.Unix(1<<62, 0))
}

// newAuthTokenExpires creates a new AuthToken that expires at the specified
// date.
func newAuthTokenExpires(user models.User, expires time.Time) (models.AuthToken, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return models.AuthToken{}, err
	}
	return models.AuthToken{
		User:    user,
		Token:   base64.RawURLEncoding.EncodeToString(token),
		Expires: expires,
	}, nil
}
