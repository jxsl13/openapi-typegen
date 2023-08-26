package strutils

import "regexp"

var (
	whitespaceRegexp = regexp.MustCompile(`\s+`)
)

// Squeeze removes all whitespace from a string.
// A b C -> AbC
func Squeeze(s string) string {
	return whitespaceRegexp.ReplaceAllString(s, "")
}
