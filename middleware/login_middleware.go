package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/common"
	"github.com/project-domino/project-domino/models"
)

// LoginMiddleware adds a user struct to the request context based on the
// authentication token provided in a cookie. Also sets a loggedIn boolean.
func LoginMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Acquire db handle from request context.
	db := context.Get(r, "db").(*gorm.DB)

	// Check if auth cookie is present.
	authCookie, err := r.Cookie("auth")

	// Go to next handler if no cookie is found
	// Set loggedIn to false
	if err == http.ErrNoCookie {
		notLoggedIn(w, r, next)
		return
	} else if err != nil {
		common.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	// If the cookie is present, search the database for the token.
	var authEntries []models.AuthToken

	db.Limit(1).Preload("User").Where(&models.AuthToken{
		Token: authCookie.Value,
	}).Where("Expires > ?", time.Now()).Find(&authEntries)
	if len(authEntries) == 0 {
		// clear the invalid/expired authtoken.
		http.SetCookie(w, &http.Cookie{
			Name:    "auth",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		notLoggedIn(w, r, next)
		return
	}

	// If there is a token and it is not expired, add user to context.
	auth := authEntries[0]
	context.Set(r, "requestUser", auth.User)
	context.Set(r, "loggedIn", true)
	next(w, r)
}

func notLoggedIn(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "requestUser", models.User{})
	context.Set(r, "loggedIn", false)
	next(w, r)
}
