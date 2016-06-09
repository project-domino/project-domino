package view

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/project-domino/project-domino/common"
)

// HomeHandler serves the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	context.Set(r, "stylesheets", []struct{ Link string }{
		{"/assets/home.css"},
	})
	common.ExecuteTemplate(w, r, "home.html")
}
