package view

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// HomeHandler serves the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
