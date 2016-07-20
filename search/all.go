package search

import "github.com/project-domino/project-domino/models"

// AllResponse holds the objects returned by a search query for all
// items
type AllResponse struct {
	Notes       []models.Note
	Collections []models.Collection
	Users       []models.User
	Tags        []models.Tag
}

// All returns a struct containing a search result for all items
func All(q string, items uint) (AllResponse, error) {
	var response AllResponse
	var searchErr error

	response.Notes, searchErr = Notes(q, 1, items)
	response.Collections, searchErr = Collections(q, 1, items)
	response.Users, searchErr = Users(q, 1, items)
	response.Tags, searchErr = Tags(q, 1, items)

	return response, searchErr
}
