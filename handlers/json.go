package handlers

import "github.com/gin-gonic/gin"

// JSON returns the value with the specified key in the request context
func JSON(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Acquire object
		v := c.MustGet(key)

		// Return object as json
		c.JSON(200, v)
	}
}
