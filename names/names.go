package names

import (
	"regexp"
	"strings"
)

var (
	nonAlphaNum   *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9-_]`)
	numericPrefix *regexp.Regexp = regexp.MustCompile(`^[0-1]+`)
)

// ToTypeName removes all special characters
// Because OperationIDs are supposed to be universally unique,
// we do not add anything in front or at the end (for now).
// Tho, identifiers starting with an integer will get a N prefix.
func ToTypeName(name string) string {
	if numericPrefix.MatchString(name) {
		name = "N" + name
	}
	return nonAlphaNum.ReplaceAllString(name, "")
}

func ToTitle(name string) string {
	return strings.Title(name)
}

// Join concatenates all strings removing duplicate overlaps between joins
// and overlaps across all previous joined strings with the next
func Join(names ...string) string {
	if len(names) == 0 {
		return ""
	} else if len(names) <= 1 {
		return names[0]
	}

	var (
		result = names[0]
		curr   string
	)
	for i := 1; i < len(names); i++ {
		curr = names[i]
		result = Merge(result, curr)
	}
	return result
}

// a and b are ordered and merged with a's suffix and b's prefix
// to deduplicate names
func Merge(a, b string) string {
	o := Overlap(a, b)
	if o > 0 {
		return a + b[o:]
	}
	return a + b
}
