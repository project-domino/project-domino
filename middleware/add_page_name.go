package middleware

import "github.com/gin-gonic/gin"

// AddPageName adds a name to the request context
func AddPageName(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pageName", name)

		c.Next()
	}
}
