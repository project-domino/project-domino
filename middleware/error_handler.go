package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/handlers"
)

// ErrorHandler is a middleware that handles errors.
//
// Errors should be used as the following:
//     _, err := functionCall()
//     if err != nil {
//         c.AbortWithError(500, errors.New("whatever error"))
//         return
//     }
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var status = 500
			if err, ok := c.Errors.Last().Err.(*errors.Error); ok {
				status = err.Status
			}
			if c.Writer.Written() {
				status = c.Writer.Status()
			}

			handlers.RenderStatusData(c, status, "error.html", "errors", struct {
				Errors     []string
				Status     int
				StatusText string
			}{c.Errors.Errors(), status, http.StatusText(status)})
		}
	}
}
