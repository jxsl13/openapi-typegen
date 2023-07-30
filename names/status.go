package names

import (
	"net/http"
	"strconv"
	"strings"
)

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
			return ToTitleTypeName(status)
		}
	}

	text := http.StatusText(int(code))
	if text == "" {
		return ToTitleTypeName(status)
	}

	text = ToTitleTypeName(strings.ToLower(text))
	return text
}

func StatusCodes(statuses []string) []string {
	result := make([]string, 0, len(statuses))
	for _, status := range statuses {
		result = append(result, StatusCode(status))
	}
	return result
}
