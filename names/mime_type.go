package names

import (
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func MimeType(mt string) string {
	m := mimetype.Lookup(mt)
	if m != nil {
		mt = m.String()
	}

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
		mt = parts[len(parts)-1]
	}

	return mt
}

func MimeTypes(mimes []string) []string {
	result := make([]string, 0, len(mimes))
	for _, mime := range mimes {
		result = append(result, MimeType(mime))
	}
	return result
}
