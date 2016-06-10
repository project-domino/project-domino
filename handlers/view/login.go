package view

import (
	"net/http"

	"github.com/project-domino/project-domino/common"
)

// LoginHandler serves the loginpage
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "login.html")
}
