package middleware

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/search"
)

// LoadSearchVars adds links to a search request for "nextPage" and "prevPage"
// It also adds the item count, query and title
func LoadSearchVars() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request variables
		q := c.DefaultQuery("q", "")
		i := c.DefaultQuery("items", "25")
		p := c.DefaultQuery("page", "1")
		searchType := c.Param("searchType")
		searchResult := c.MustGet("searchResult")

		// Convert page and items to uint
		tItems, convertErr1 := strconv.ParseInt(i, 10, 64)
		tPage, convertErr2 := strconv.ParseInt(p, 10, 64)
		if convertErr1 != nil || convertErr2 != nil {
			errors.BadParameters.Apply(c)
			return
		}
		items := int(tItems)
		page := int(tPage)

		// Set query and items to request context
		c.Set("items", items)
		c.Set("query", q)

		// Set nextPage and prevPage to request context
		var nextPage string
		var prevPage string
		currentPage := fmt.Sprintf("/search/%s?q=%s&items=%d&page=%d",
			searchType, q, items, page)
		if searchType == search.AllSearchType {
			nextPage = "#"
			prevPage = "#"
		} else {
			if reflect.ValueOf(searchResult).Len() <= items || page == search.MaxPage {
				nextPage = currentPage
			} else {
				nextPage = fmt.Sprintf("/search/%s?q=%s&items=%d&page=%d",
					searchType, q, items, page+1)
			}

			if page > 1 {
				prevPage = fmt.Sprintf("/search/%s?q=%s&items=%d&page=%d",
					searchType, q, items, page-1)
			} else {
				prevPage = currentPage
			}
		}

		c.Set("nextPage", nextPage)
		c.Set("prevPage", prevPage)

		c.Next()
	}
}
