package strutils

import (
	"regexp"
	"strings"
)

var (
	nonAlphaNum *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

// TypeName removes all non alpha numeric characters
// Integers are prefixed with an "N"
// The first character is capitalized.
func TypeName(name string) string {
	name = PrefixInteger(name)
	return ToTitle(toTypeName(name))
}

// MergeAllTypeNames merges all type names and prefixes the result with an "N" if the result starts with an integer.
func MergeAllTypeNames(names ...string) string {
	titledTypes := make([]string, 0, len(names))
	for _, name := range names {
		titledTypes = append(titledTypes, toTypeName(ToTitle(name)))
	}

	return PrefixInteger(MergeAll(titledTypes...))
}

func toTypeName(name string) string {
	// some api specs use weird names
	// we make plural from array types
	if strings.HasSuffix(name, "[]") {
		name = MergeAll(name[:len(name)-2], "s")
	}

	name = UnwrapPathParameters(ToTitle, name)

	return nonAlphaNum.ReplaceAllString(name, "")
}
