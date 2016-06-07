package view

import (
	"net/http"

	"bitbucket.org/akshay_shekhawat/project-domino/common"
)

// SubjectListHandler serves the list of subjects
func SubjectListHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
