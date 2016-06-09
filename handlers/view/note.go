package view

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// NoteHandler serves an individual note
func NoteHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
