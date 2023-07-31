package strutils

import (
	"net/http"
	"strconv"
	"strings"
)

// StatusCode returns the status code name for a given status code.
// If the status code is not a valid integer, it will be returned as is.
// If the status code is a valid integer, but not a valid http status code,
// it will be returned as is.
func StatusCode(status string) string {
	code, err := strconv.ParseInt(status, 10, 32)
	if err != nil {
		switch strings.ToUpper(status) {
		case "1XX":
			return "Info"
		case "2XX":
			return "Success"
		case "3XX":
			return "Redirect"
		case "4XX":
			return "ClientError"
		case "5XX":
			return "ServerError"
		default:
			return TypeName(status)
		}
	}

	text := http.StatusText(int(code))
	if text == "" {
		return TypeName(status)
	}

	text = TypeName(strings.ToLower(text))
	return text
}

func StatusCodes(statuses []string) []string {
	result := make([]string, 0, len(statuses))
	for _, status := range statuses {
		result = append(result, StatusCode(status))
	}
	return result
}
