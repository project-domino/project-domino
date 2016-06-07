package api

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino-server/common"
	"github.com/project-domino/project-domino-server/models"
)

// LoginHandler handles requests to log a user in.
// If credentials are valid, sets an auth cookie.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get needed variables from request.
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// If there are blank fields, return bad request
	if password == "" || (userName == "" && email == "") {
		common.HandleError(w, errors.New("Missing Parameters"), http.StatusBadRequest)
		return
	}
	// If both userName and email are filled in, return bad request
	if userName != "" && email != "" {
		common.HandleError(w, errors.New("Both username and email cannot be used."), http.StatusBadRequest)
		return
	}

	// Acquire db handle from request context.
	db := context.Get(r, "db").(*gorm.DB)

	// Find user in the database
	var users []models.User
	db.Limit(1).
		Where(&models.User{
			Email: email,
		}).Or(&models.User{
		UserName: userName,
	}).Find(&users)

	// If a user with these credentials does not exist, return error
	if len(users) == 0 {
		common.HandleError(w, errors.New("Invalid Credentials"), http.StatusForbidden)
		return
	}

	// Otherwise, check password and assign cookie
	user := users[0]
	if !user.CheckPassword(password) {
		common.HandleError(w, errors.New("Invalid Credentials"), http.StatusForbidden)
		return
	}
	AuthCookie(w, r, user, db)
}

// LogoutHandler handles requests to log a user out.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

// RegisterHandler handles requests to handle a new user.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Get needed variables from request.
	email := r.FormValue("email")
	userName := r.FormValue("userName")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	password := r.FormValue("password")
	retypePassword := r.FormValue("retypePassword")

	// Check if the request is missing needed parameters
	if userName == "" || password == "" || retypePassword == "" {
		common.HandleError(w, errors.New("Missing Parameters"), http.StatusBadRequest)
		return
	}

	// Check if password matches retype password
	if retypePassword != password {
		common.HandleError(w, errors.New("Passwords do not match."), http.StatusBadRequest)
		return
	}

	// Acquire db handle from request context.
	db := context.Get(r, "db").(*gorm.DB)

	// Check if other users have the same email or userName
	var checkUsers []models.User
	db.Where(&models.User{
		Email: email,
	}).Or(&models.User{
		UserName: userName,
	}).Find(&checkUsers)
	if len(checkUsers) != 0 {
		common.HandleError(w, errors.New("User with same email/username already exists."), http.StatusForbidden)
		return
	}

	// Create the user.
	user := models.User{
		Email:     email,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Type:      models.General,
	}
	if err := user.SetPassword(password); err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	// Add user to database.
	db.Create(user)

	// Set an auth cookie for the user
	AuthCookie(w, r, user, db)
	http.Redirect(w, r, "/", http.StatusFound)
}

// AuthCookie creates an authentication token and sends it to the client.
func AuthCookie(w http.ResponseWriter, r *http.Request, user models.User, db *gorm.DB) {
	var authToken models.AuthToken
	var err error

	if strings.Contains(r.UserAgent(), "Domino") {
		authToken, err = newAuthToken(user)
	} else {
		authToken, err = newAuthTokenExpires(user, time.Now().Add(time.Hour*24*7))
	}

	if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	db.Create(&authToken)

	// Create cookie
	cookie := http.Cookie{
		Name:    "auth",
		Value:   authToken.Token,
		Expires: authToken.Expires,
	}

	http.SetCookie(w, &cookie)
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
