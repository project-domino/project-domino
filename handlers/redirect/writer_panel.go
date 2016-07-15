package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WriterPanel redirects the user to a page in the writer panel
// TODO redirect to latest draft
func WriterPanel(c *gin.Context) {
	c.Redirect(http.StatusFound, "/writer-panel/note")
}
