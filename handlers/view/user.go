package view

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// UserHandler serves a webpage for an individual user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
