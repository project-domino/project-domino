package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
)

// MaxItems is the maximum number of items that can be returned
const MaxItems int = 100

// MaxPage is the maximum page number which can be returned
const MaxPage int = 10

// getListVars returns a tuple which contains (items, page, err)
func getListVars(c *gin.Context) (int, int, *errors.Error) {
	// Get values from request context
	i := c.DefaultQuery("items", "25")
	p := c.DefaultQuery("page", "1")

	// Convert page and items to int
	tItems, convertErr1 := strconv.ParseInt(i, 10, 64)
	tPage, convertErr2 := strconv.ParseInt(p, 10, 64)
	items := int(tItems)
	page := int(tPage)

	// Verify valid parameters
	if convertErr1 != nil || items <= 0 || items > MaxItems {
		return items, page, errors.InvalidItems
	}
	if convertErr2 != nil || page <= 0 || page > MaxPage {
		return items, page, errors.InvalidPage
	}

	return items, page, nil
}

// getPages returns the next and prev pages
func getPages(currentPage int, isMax bool) (int, int) {
	var nextPage int
	var prevPage int

	if isMax || currentPage == MaxPage {
		nextPage = currentPage
	} else {
		nextPage = currentPage + 1
	}

	if currentPage > 1 {
		prevPage = currentPage - 1
	} else {
		prevPage = currentPage
	}

	return prevPage, nextPage
}
