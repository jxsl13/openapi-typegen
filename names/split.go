package names

import "github.com/fatih/camelcase"

// Split splits a string into a slice of words by splitting on camelcase.
func Split(s string) []string {
	return camelcase.Split(s)
}

// SplitAll splits a slice of strings into a slice of words by splitting on camelcase.
func SplitAll(names ...string) []string {
	ss := make([]string, 0, len(names))
	for _, name := range names {
		ss = append(ss, Split(name)...)
	}
	return ss
}
