package redirect

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/models"
)

// User redirects the user to a different user page
func User(c *gin.Context) {
	username := c.Param("username")
	user := c.MustGet("pageUser").(models.User)

	var linkFormat string

	if user.Type == "writer" || user.Type == "admin" {
		linkFormat = "/u/%s/notes"
	} else {
		linkFormat = "/u/%s/upvote-notes"
	}

	c.Redirect(http.StatusFound, fmt.Sprintf(linkFormat, username))
}
