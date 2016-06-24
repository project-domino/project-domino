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

			c.Negotiate(errCode, gin.Negotiate{
				Offered: []string{
					gin.MIMEHTML,
					gin.MIMEXML,
					gin.MIMEJSON,
				},
				HTMLName: "error.html",
				Data:     c.Errors,
			})
		}
	}
}
