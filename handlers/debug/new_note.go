package debug

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// NewNoteHandler returns the page for creating a new note
func NewNoteHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "new-note.html")
}
