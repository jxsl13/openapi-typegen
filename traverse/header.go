package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Header traverses the given header and all unique non-reference schemas in it.
func Header(header *openapi3.HeaderRef, visitor SchemaVisitor, levelNames map[string][]string) error {
	if header == nil {
		return nil
	}
	if header.Ref != "" {
		return nil
	}
	if header.Value == nil {
		return nil
	}

	// we want to handle component header definitions like any other parameter
	return ParameterSchema(&header.Value.Parameter, visitor, add(levelNames, InKey, openapi3.ParameterInHeader))
}
