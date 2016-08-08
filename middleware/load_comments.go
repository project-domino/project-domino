package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadComments loads comments into the request context with specified objects
// If :commentType is present in the URL, that comment type will be loaded
// Otherwise Questions will be loaded
func LoadComments(objects ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get values from request context
		i := c.DefaultQuery("items", "25")
		p := c.DefaultQuery("page", "1")
		noteID := c.Param("noteID")
		commentType := c.Param("commentType")

		// Acquire user from the request context.
		user := c.MustGet("user").(models.User)

		// Verify comment type is valid
		if (commentType != models.QuestionComment) && (commentType != models.SuggestionComment) {
			commentType = models.QuestionComment
		}

		// Convert page and items to int
		tItems, convertErr1 := strconv.ParseInt(i, 10, 64)
		tPage, convertErr2 := strconv.ParseInt(p, 10, 64)
		items := int(tItems)
		page := int(tPage)

		// Verify valid parameters
		if convertErr1 != nil || items <= 0 {
			errors.InvalidItems.Apply(c)
			return
		}
		if convertErr2 != nil || page <= 0 {
			errors.InvalidPage.Apply(c)
			return
		}

		// Set objects to be preloaded to db
		preloadedDB := db.DB
		for _, object := range objects {
			preloadedDB = preloadedDB.Preload(object)
		}

		// If there is a user get the users comments first
		if user.ID != 0 {
			preloadedDB = preloadedDB.Order(fmt.Sprintf("(user_id = %d)", user.ID))
		}

		// Query for parent comments
		var parentComments []models.Comment
		var parentIDs []uint
		if err := preloadedDB.
			Where("note_id = ?", noteID).
			Where("parent_id = ?", 0).
			Where("type = ?", commentType).
			Order("created_at desc").
			Limit(items).
			Offset((page-1)*items).
			Find(&parentComments).
			Pluck("parent_id", &parentIDs).
			Error; err != nil {

			c.AbortWithError(500, err)
			return
		}

		// Get childComments
		var childComments []models.Comment
		if err := preloadedDB.
			Where("note_id = ?", noteID).
			Where("parent_id IN (?)", parentIDs).
			Where("type = ?", commentType).
			Order("created_at desc").
			Find(&parentComments).
			Error; err != nil {

			c.AbortWithError(500, err)
			return
		}

		// Append child comments to parent
		for _, child := range childComments {
			for _, parent := range parentComments {
				if parent.ID == child.ParentID {
					parent.Children = append(parent.Children, child)
				}
			}
		}

		// Add comments to request context
		c.Set("comments", parentComments)

		c.Next()
	}
}
