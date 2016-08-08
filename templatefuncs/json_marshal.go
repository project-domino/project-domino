package templatefuncs

import (
	"encoding/json"
	"html/template"
)

// JSONMarshal turns an object into json for templates
func JSONMarshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}
