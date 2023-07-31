package strutils

import (
	"strings"
)

// Append appends the given strings to the first string.
func Append(str string, strs ...string) string {
	return strings.Join(append([]string{str}, strs...), "")
}

// Ignore returns all strings that are not equal to the ignored string.
// Equality check is case insensitive.
func Ignore(ignored string, names ...string) []string {
	result := make([]string, 0, len(names))
	for _, name := range names {
		if !strings.EqualFold(ignored, name) {
			result = append(result, name)
		}
	}
	return result
}
