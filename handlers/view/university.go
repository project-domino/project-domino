package view

import (
	"net/http"

	"bitbucket.org/akshay_shekhawat/project-domino/common"
)

// UniversityHandler serves the webpage for an individual university
func UniversityHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
