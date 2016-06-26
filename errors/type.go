package errors

import "github.com/gin-gonic/gin"

// Error is an error struct used for most errors encountered.
type Error struct {
	Status  int
	Message string
}

// Apply adds the error to the current context.
func (e *Error) Apply(c *gin.Context) {
	c.AbortWithError(e.Status, e)
}

func (e *Error) Error() string {
	return e.Message
}
