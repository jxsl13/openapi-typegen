package traverse

import "github.com/getkin/kin-openapi/openapi3"

// RequestBody traverses the given request body and all unique non-reference schemas in it.
func RequestBody(requestBody *openapi3.RequestBodyRef, visitor SchemaVisitor, levelNames map[string][]string) error {
	if requestBody == nil {
		return nil
	}
	if requestBody.Ref != "" {
		return nil
	}
	if requestBody.Value == nil {
		return nil
	}

	var err error
	for contentType, mediaType := range requestBody.Value.Content {
		if mediaType == nil {
			continue
		}
		err = MediaType(mediaType, visitor, add(levelNames, ContentKey, contentType, TypeKey, RequestType))
		if err != nil {
			return err
		}
	}
	return nil
}
