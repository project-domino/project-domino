package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/handlers/vars"
)

// Simple creates a new handler that renders based on the given template name.
func Simple(template string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, template, vars.Default(c))
	}
}
