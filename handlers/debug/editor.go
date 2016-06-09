package debug

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// EditorHandler serves the not-yet-complete editor.
func EditorHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "editor.html")
}
