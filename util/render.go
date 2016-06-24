package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/handlers/vars"
)

// Render acts as RenderStatusData with a status of http.StatusOK and data
// loaded from the context with the key dataName.
func Render(c *gin.Context, htmlPage string, dataName string) {
	RenderStatus(c, http.StatusOK, htmlPage, dataName)
}

// RenderStatus acts as RenderStatusData with data loaded from the context with
// the key dataName.
func RenderStatus(c *gin.Context, status int, htmlPage string, dataName string) {
	value, ok := c.Get(dataName)
	if !ok {
		value = nil
	}
	RenderStatusData(c, status, htmlPage, dataName, value)
}

// RenderData acts as RenderStatusData with a status of http.StatusOK.
func RenderData(c *gin.Context, htmlPage string, dataName string, dataValue interface{}) {
	RenderStatusData(c, http.StatusOK, htmlPage, dataName, dataValue)
}

// RenderStatusData sends the data with the given status, rendering as JSON, a
// plain-text string, XML, or inserts it into an HTML page with the given
// variable name. The other variables used in the HTML page rendering are the
// same as in the "github.com/project-domino/project-domino/handlers/vars"
// package.
func RenderStatusData(c *gin.Context, status int, htmlPage string, dataName string, dataValue interface{}) {
	switch c.NegotiateFormat(
		gin.MIMEHTML,
		gin.MIMEJSON,
		gin.MIMEPlain,
		gin.MIMEXML,
	) {
	case gin.MIMEHTML:
		data := vars.Default(c)
		if dataName != "" {
			data[dataName] = dataValue
		}
		c.HTML(status, htmlPage, data)
	case gin.MIMEJSON:
		c.JSON(status, dataValue)
	case gin.MIMEPlain:
		c.String(status, "%s", dataValue)
	case gin.MIMEXML:
		c.XML(status, dataValue)
	}
}
