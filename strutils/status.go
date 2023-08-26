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
		return RangeStatusCode(status)
	}

	text := http.StatusText(int(code))
	if text != "" {
		return TypeName(Squeeze(text))
	}
	return RangeStatusCode(status)
}

// RangeStatusCode converts nXX status codes to a string representation.
// 1XX -> Info
// 2XX -> Success
// 3XX -> Redirect
// 4XX -> ClientError
// 5XX -> ServerError
// Any other will be converted to a type name.
func RangeStatusCode(nxx string) string {
	switch strings.ToUpper(nxx) {
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
		return TypeName(nxx)
	}
}

func StatusCodes(statuses []string) []string {
	result := make([]string, 0, len(statuses))
	for _, status := range statuses {
		result = append(result, StatusCode(status))
	}
	return result
}
