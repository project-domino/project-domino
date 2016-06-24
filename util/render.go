package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/handlers/vars"
)

// Render acts as RenderStatusData with a status of http.StatusOK and nil data.
func Render(c *gin.Context, htmlPage string) {
	RenderStatusData(c, http.StatusOK, htmlPage, nil)
}

// RenderStatus acts as RenderStatusData with nil data.
func RenderStatus(c *gin.Context, status int, htmlPage string) {
	RenderStatusData(c, status, htmlPage, nil)
}

// RenderData acts as RenderStatusData with a status of http.StatusOK.
func RenderData(c *gin.Context, htmlPage string, data interface{}) {
	RenderStatusData(c, http.StatusOK, htmlPage, data)
}

// RenderStatusData sends the data with the given status, rendering as JSON, a
// plain-text string, XML, or inserts it into an HTML page as the "data"
// variable. The other variables used in the HTML page rendering are the same as
// in the "github.com/project-domino/project-domino/handlers/vars" package.
func RenderStatusData(c *gin.Context, status int, htmlPage string, data interface{}) {
	switch c.NegotiateFormat(
		gin.MIMEHTML,
		gin.MIMEJSON,
		gin.MIMEPlain,
		gin.MIMEXML,
	) {
	case gin.MIMEHTML:
		c.HTML(status, htmlPage, vars.Default(c).Set("data", data))
	case gin.MIMEJSON:
		c.JSON(status, data)
	case gin.MIMEPlain:
		c.String(status, "%s", data)
	case gin.MIMEXML:
		c.XML(status, data)
	}
}
