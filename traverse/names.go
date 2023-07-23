package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
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
