package redirect

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User redirects the user to a different user page
func User(c *gin.Context) {
	username := c.Param("username")

	c.Redirect(http.StatusFound, fmt.Sprintf("/u/%s/notes", username))
}
