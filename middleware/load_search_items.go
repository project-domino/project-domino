package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/search"
)

// LoadSearchItems loads items that match a given query into the request context
func LoadSearchItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request variables
		q := c.DefaultQuery("q", "")
		i := c.DefaultQuery("items", "25")
		p := c.DefaultQuery("page", "1")
		searchType := c.Param("searchType")

		// Convert page and items to uint
		tempItems, convertErr1 := strconv.ParseUint(i, 10, 64)
		tempPage, convertErr2 := strconv.ParseUint(p, 10, 64)
		if convertErr1 != nil || convertErr2 != nil {
			errors.BadParameters.Apply(c)
			return
		}
		items := uint(tempItems)
		page := uint(tempPage)

		// Verify item and page numbers
		if items <= 0 || items > search.MaxItems {
			errors.InvalidItems.Apply(c)
			return
		}
		if page <= 0 || page > search.MaxPage {
			errors.InvalidPage.Apply(c)
			return
		}

		// Search db for objects
		var results interface{}
		var err error

		switch searchType {
		case search.AllSearchType:
			results, err = search.All(q, items)
		case search.NoteSearchType:
			results, err = search.Notes(q, items, page)
		case search.CollectionSearchType:
			results, err = search.Collections(q, items, page)
		case search.UserSearchType:
			results, err = search.Users(q, items, page)
		case search.TagSearchType:
			results, err = search.Tags(q, items, page)
		default:
			errors.NotFound.Apply(c)
			return
		}

		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.Set("searchResult", results)

		c.Next()
	}
}
