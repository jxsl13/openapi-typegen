package strutils

import (
	"strings"
)

// MimeTypeName returns the last part of a mime type.
// And converts the mime type to a Go type name.
func MimeTypeName(mt string) string {
	parts := strings.Split(mt, ";")
	if len(parts) == 2 {
		mt = parts[0]
	}

	parts = strings.Split(mt, "/")
	if len(parts) == 2 {
		mt = parts[len(parts)-1]
	}

	parts = strings.Split(mt, "+")
	if len(parts) == 2 {
		return MergeAllTypeNames(parts...)
	}

	return MergeAllTypeNames(mt)
}

// MimeTypeNames takes the mime types and returns the last part of each mime type.
// And converts the mime types to Go type names.
func MimeTypeNames(mimes []string) []string {
	result := make([]string, 0, len(mimes))
	for _, mime := range mimes {
		result = append(result, MimeTypeName(mime))
	}
	return result
}
