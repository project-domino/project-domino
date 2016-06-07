package view

import (
	"net/http"

	"bitbucket.org/akshay_shekhawat/project-domino/common"
)

// UserHandler serves a webpage for an individual user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
