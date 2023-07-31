package strutils

import "regexp"

var (
	numericPrefix *regexp.Regexp = regexp.MustCompile(`^[0-9]+`)
)

// PrefixInteger prefixes the given name with an "N" if it is a valid integer.
func PrefixInteger(name string) string {
	if numericPrefix.MatchString(name) {
		name = "N" + name
	}
	return name
}
