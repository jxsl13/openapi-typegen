package traverse

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

var (
	RequestBodyNameSuffix = "Request"
	ResponseNameSuffix    = "Response"
	ParameterNameSuffix   = "Parameter"
	HeaderNameSuffix      = "Header"
)

func NameFromOperation(method, path string, operation *openapi3.Operation, suffix string) string {
	if operation.OperationID != "" {
		return names.Join(
			names.ToTitleTypeName(operation.OperationID),
			suffix,
		)
	}
	return names.Join(
		names.ToTitle(method),
		names.ToTitleTypeName(path),
		suffix,
	)
}

func NameFromOperationParameterRef(method, path, operationID string, parameter *openapi3.ParameterRef, suffix string) string {
	if operationID != "" {
		return names.Join(
			names.ToTitleTypeName(operationID),
			names.ToTitle(parameter.Value.In),
			names.ToTitleTypeName(parameter.Value.Name),
			suffix,
		)
	}
	return names.Join(
		names.ToTitle(method),
		names.ToTitleTypeName(path),
		names.ToTitle(parameter.Value.In),
		names.ToTitleTypeName(parameter.Value.Name),
		suffix,
	)
}

// NameFromParameter is used when deriving a parameter name from a parameter defined direcly below
// a path.
func NameFromParameter(path string, parameter *openapi3.Parameter, suffix string) string {
	return names.Join(
		names.ToTitleTypeName(path),
		names.ToTitle(parameter.In),
		names.ToTitleTypeName(parameter.Name),
		suffix,
	)
}

func NameFromComponentMediaType(name, mimeType string, suffix string) string {
	return names.Join(
		names.ToTitleTypeName(name),
		names.ToTitleTypeName(mimeType),
		names.ToTitleTypeName(suffix),
	)
}

func NameFromStatusCode(status string) string {
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
			return names.ToTitleTypeName(status)
		}
	}

	return names.ToTitleTypeName(http.StatusText(int(code)))
}
