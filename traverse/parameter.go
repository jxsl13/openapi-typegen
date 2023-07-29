package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Parameter traverses the given parameter and all unique non-reference schemas in it.
func Parameter(parameter *openapi3.ParameterRef, visitor SchemaVisitor, levelNames map[string][]string) error {
	if parameter == nil {
		return nil
	}
	if parameter.Ref != "" {
		return nil
	}
	if parameter.Value == nil {
		return nil
	}
	err := ParameterSchema(parameter.Value, visitor, levelNames)
	if err != nil {
		return err
	}
	return nil
}

func ParameterSchema(parameter *openapi3.Parameter, visitor SchemaVisitor, levelNames map[string][]string) error {
	if parameter == nil {
		return nil
	}
	if parameter.Schema == nil {
		return nil
	}

	// for headers the IN value can become an empty string
	return visitor(parameter.Schema, add(levelNames, NameKey, parameter.Name, InKey, parameter.In, TypeKey, ParameterType))
}
