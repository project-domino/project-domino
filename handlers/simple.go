package handlers

import "github.com/gin-gonic/gin"

// Simple creates a new handler that renders based on the given template name.
func Simple(template string) gin.HandlerFunc {
	return func(c *gin.Context) {
		Render(c, template, "")
	}
}

// Value creates a new handler that renders a page, with the given value from
// the context attached.
func Value(template, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		Render(c, template, key)
	}
}
