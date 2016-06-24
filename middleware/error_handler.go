package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
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
			var errCode = 500
			switch err := c.Errors[0]; err {
			default:
				log.Printf("Unknown error: %T %v", err, err)
			}
			if c.Writer.Written() {
				errCode = c.Writer.Status()
			}

			switch c.NegotiateFormat(
				gin.MIMEHTML,
				gin.MIMEJSON,
				gin.MIMEPlain,
				gin.MIMEXML,
			) {
			case gin.MIMEHTML:
				c.HTML(errCode, "error.html", c.Errors)
			case gin.MIMEJSON:
				c.JSON(errCode, c.Errors)
			case gin.MIMEPlain:
				c.String(errCode, "%s", c.Errors)
			case gin.MIMEXML:
				c.XML(errCode, c.Errors)
			}
		}
	}
}
