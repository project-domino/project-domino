package view

import (
	"net/http"

	"github.com/project-domino/project-domino-server/common"
)

// SubjectListHandler serves the list of subjects
func SubjectListHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
