package vars

import "github.com/gin-gonic/gin"

// Default returns the template variables that are available to all templates.
// At the current time, this includes those returned by Keys and Params.
func Default(c *gin.Context) Vars {
	return New(
		Keys(c),
		Params(c),
	)
}
