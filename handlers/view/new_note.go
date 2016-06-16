package view

import (
	"errors"
	"net/http"

	"github.com/project-domino/project-domino/models"

	"github.com/gorilla/context"
	"github.com/project-domino/project-domino/common"
)

// NewNoteHandler returns the page for creating a new note
func NewNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Verify the user making the request is logged in
	loggedIn := context.Get(r, "loggedIn").(bool)
	if !loggedIn {
		common.HandleError(w, errors.New("You must be logged in to perform this action."), http.StatusForbidden)
		return
	}

	// Check if the user is a writer or admin
	requestUser := context.Get(r, "requestUser").(models.User)
	if !(requestUser.Type == models.Admin || requestUser.Type == models.Writer) {
		common.HandleError(w, errors.New("You do not have access to this feature."), http.StatusForbidden)
		return
	}

	common.ExecuteTemplate(w, r, "new-note.html")
}
