package view

import (
	"net/http"

	"github.com/project-domino/project-domino-server/common"
)

// SearchHandler returns a view for the search page
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
