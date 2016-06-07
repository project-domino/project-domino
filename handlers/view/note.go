package view

import (
	"net/http"

	"bitbucket.org/akshay_shekhawat/project-domino/common"
)

// NoteHandler serves an individual note
func NoteHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
