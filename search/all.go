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
func All(q string, items int) (AllResponse, error) {
	var response AllResponse
	var searchErr error

	response.Notes, searchErr = Notes(q, items, 1)
	response.Collections, searchErr = Collections(q, items, 1)
	response.Users, searchErr = Users(q, items, 1)
	response.Tags, searchErr = Tags(q, items, 1)

	return response, searchErr
}
