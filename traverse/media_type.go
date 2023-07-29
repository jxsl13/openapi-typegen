package traverse

import "github.com/getkin/kin-openapi/openapi3"

// MediaType traverses the given media type and all unique non-reference schemas in it.
func MediaType(mediaType *openapi3.MediaType, visitor SchemaVisitor, levelNames map[string][]string) error {
	if mediaType == nil {
		return nil
	}
	if mediaType.Schema == nil {
		return nil
	}

	if mediaType.Schema.Ref != "" {
		return nil
	}

	return visitor(mediaType.Schema, levelNames)
}
