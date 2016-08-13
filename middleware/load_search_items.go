package middleware

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/search"
)

// LoadSearchItems loads items that match a given query into the request context
func LoadSearchItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request variables
		q := c.DefaultQuery("q", "")
		searchType := c.Param("searchType")
		items, page, err := getListVars(c)
		if err != nil {
			err.Apply(c)
			return
		}

		// Set query and items to request context
		c.Set("items", items)
		c.Set("query", q)

		// Search db for objects
		var results interface{}
		var searchErr error

		switch searchType {
		case search.AllSearchType:
			results, searchErr = search.All(q, 5)
		case search.NoteSearchType:
			results, searchErr = search.Notes(q, items, page)
		case search.CollectionSearchType:
			results, searchErr = search.Collections(q, items, page)
		case search.UserSearchType:
			results, searchErr = search.Users(q, items, page)
		case search.TagSearchType:
			results, searchErr = search.Tags(q, items, page)
		default:
			errors.NotFound.Apply(c)
			return
		}

		if searchErr != nil {
			errors.DB.Apply(c)
			return
		}

		// Set current, next and prev pages
		if searchType != search.AllSearchType {
			prevPage, nextPage := getPages(page, reflect.ValueOf(results).Len() < items)
			c.Set("currentPage", page)
			c.Set("nextPage", nextPage)
			c.Set("prevPage", prevPage)
		}

		c.Set("searchResult", results)

		c.Next()
	}
}
