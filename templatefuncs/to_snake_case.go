package templatefuncs

import "regexp"

var snakeCaseRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

// ToSnakeCase converts a string to snake_case
// This is primarily used to create readable links
func ToSnakeCase(s string) string {
	return snakeCaseRegex.ReplaceAllString(s, "_")
}
