package view

import (
	"net/http"

	"bitbucket.org/akshay_shekhawat/project-domino/common"
)

// TopicHandler serves a list of notes in a given topic
func TopicHandler(w http.ResponseWriter, r *http.Request) {
	common.ExecuteTemplate(w, r, "home.html")
}
