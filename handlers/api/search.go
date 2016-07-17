package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
	"github.com/project-domino/project-domino/search"
)

// SearchResponse holds the response object for a search query
type SearchResponse struct {
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

// Search handles searching the db
func Search(c *gin.Context) {
	// Get query from request
	q := c.DefaultQuery("q", "")

	// Parse search query
	searchQuery, err := search.ParseQuery(q)
	if err != nil {
		errors.BadParameters.Apply(c)
		return
	}
	querySelectors := searchQuery.Selectors
	queryText := strings.Join(searchQuery.Text, " & ")

	var response SearchResponse
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

	if searchErr != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Send back response in json
	c.JSON(http.StatusOK, response)
}

func noteSearch(q string, selectors map[string][]string) ([]models.Note, error) {
	var notes []models.Note
	if q != "" {
		db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Where("published = ?", true).
			Find(&notes)
	}
	return notes, nil
}

func tagSearch(q string, selectors map[string][]string) ([]models.Tag, error) {
	var tags []models.Tag
	if q != "" {
		db.DB.Where("searchtext @@ to_tsquery(?)", q).Find(&tags)
	}
	return tags, nil
}

func collectionSearch(q string, selectors map[string][]string) ([]models.Collection, error) {
	var collections []models.Collection
	if q != "" {
		db.DB.Where("searchtext @@ to_tsquery(?)", q).
			Where("published = ?", true).
			Find(&collections)
	}
	return collections, nil
}

func userSearch(q string, selectors map[string][]string) ([]models.User, error) {
	var users []models.User
	if q != "" {
		db.DB.Where("searchtext @@ to_tsquery(?)", q).Find(&users)
	}
	return users, nil
}

func universitySearch(q string, selectors map[string][]string) ([]models.University, error) {
	var universities []models.University
	if q != "" {
		db.DB.Where("searchtext @@ to_tsquery(?)", q).Find(&universities)
	}
	return universities, nil
}
