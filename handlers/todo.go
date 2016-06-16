package handlers

import "github.com/gin-gonic/gin"

// TODO is a handler that serves a TODO message.
func TODO(c *gin.Context) {
	c.String(200, "TODO %s", c.Request.URL.String())
}
