package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Account redirects the user to a page in the writer panel
func Account(c *gin.Context) {
	c.Redirect(http.StatusFound, "/account/profile")
}
