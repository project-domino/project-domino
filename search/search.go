package search

import (
	"strings"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Response holds the response object for a search query
type Response struct {
	Notes        []models.Note
	Collections  []models.Collection
	Users        []models.User
	Universities []models.University
	Tags         []models.Tag
}

// These constants are the valid search query values for type.
const (
	NoteSearchType       string = "note"
	CollectionSearchType        = "collection"
	UserSearchType              = "user"
	UniversitySearchType        = "university"
	TagSearchType               = "tag"
)

// Search returns a Response for a given query
func Search(q string) (Response, error) {
	var response Response

	// Parse search query
	searchQuery, err := ParseQuery(q)
	if err != nil {
		return response, err
	}
	querySelectors := searchQuery.Selectors
	queryText := strings.Join(searchQuery.Text, " & ")

	var searchErr error

	// If types are specified in a query, only search for those types
	if val, ok := querySelectors["type"]; ok {
		for _, selectorType := range val {
			switch selectorType {
			case NoteSearchType:
				response.Notes, searchErr = noteSearch(queryText, querySelectors)
			case CollectionSearchType:
				response.Collections, searchErr = collectionSearch(queryText, querySelectors)
			case UserSearchType:
				response.Users, searchErr = userSearch(queryText, querySelectors)
			case UniversitySearchType:
				response.Universities, searchErr = universitySearch(queryText, querySelectors)
			case TagSearchType:
				response.Tags, searchErr = tagSearch(queryText, querySelectors)
			}
		}
	} else {
		response.Notes, searchErr = noteSearch(queryText, querySelectors)
		response.Collections, searchErr = collectionSearch(queryText, querySelectors)
		response.Users, searchErr = userSearch(queryText, querySelectors)
		response.Universities, searchErr = universitySearch(queryText, querySelectors)
		response.Tags, searchErr = tagSearch(queryText, querySelectors)
	}

	return response, searchErr

}

func noteSearch(q string, selectors map[string][]string) ([]models.Note, error) {
	var notes []models.Note
	if q != "" {
		if err := db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Where("published = ?", true).
			Find(&notes).
			Error; err != nil {
			return notes, err
		}
	}
	return notes, nil
}

func tagSearch(q string, selectors map[string][]string) ([]models.Tag, error) {
	var tags []models.Tag
	if q != "" {
		if err := db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Find(&tags).
			Error; err != nil {
			return tags, err
		}
	}
	return tags, nil
}

func collectionSearch(q string, selectors map[string][]string) ([]models.Collection, error) {
	var collections []models.Collection
	if q != "" {
		if err := db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Where("published = ?", true).
			Find(&collections).
			Error; err != nil {
			return collections, err
		}
	}
	return collections, nil
}

func userSearch(q string, selectors map[string][]string) ([]models.User, error) {
	var users []models.User
	if q != "" {
		if err := db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Find(&users).
			Error; err != nil {
			return users, err
		}
	}
	return users, nil
}

func universitySearch(q string, selectors map[string][]string) ([]models.University, error) {
	var universities []models.University
	if q != "" {
		if err := db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Find(&universities).
			Error; err != nil {
			return universities, err
		}
	}
	return universities, nil
}
