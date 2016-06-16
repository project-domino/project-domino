package vars

import "github.com/gin-gonic/gin"

// Params returns the variables that are in the parameters of a route.
func Params(c *gin.Context) Vars {
	vars := make(Vars)
	for _, param := range c.Params {
		vars[param.Key] = param.Value
	}
	return vars
}
