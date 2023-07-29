package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Parameter traverses the given parameter and all unique non-reference schemas in it.
func Parameter(parameter *openapi3.ParameterRef, visitor SchemaVisitor, levelNames ...string) error {
	if parameter == nil {
		return nil
	}
	if parameter.Ref != "" {
		return nil
	}
	if parameter.Value == nil {
		return nil
	}
	if err := ParameterSchema(parameter.Value, visitor, append(levelNames, parameter.Value.Name, parameter.Value.In)...); err != nil {
		return err
	}
	return nil
}

func ParameterSchema(parameter *openapi3.Parameter, visitor SchemaVisitor, levelNames ...string) error {
	if parameter == nil {
		return nil
	}
	if parameter.Schema == nil {
		return nil
	}
	return visitor(parameter.Schema, append(levelNames, ParameterSuffix)...)
}
