package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Header traverses the given header and all unique non-reference schemas in it.
func Header(header *openapi3.HeaderRef, visitor SchemaVisitor, levelNames ...string) error {
	if header == nil {
		return nil
	}
	if header.Ref != "" {
		return nil
	}
	if header.Value == nil {
		return nil
	}
	if header.Value.Schema == nil {
		return nil
	}
	return visitor(header.Value.Schema, append(levelNames, HeaderSuffix)...)
}
