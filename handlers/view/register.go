package view

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// RegisterHandler serves the registerpage
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "register.html")
}
