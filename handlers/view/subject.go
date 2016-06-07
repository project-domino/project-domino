package view

import (
	"net/http"

	"github.com/project-domino/project-domino-server/common"
)

// SubjectHandler serves a list of topics in a given subject
func SubjectHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
