package traverse

import "github.com/getkin/kin-openapi/openapi3"

// RequestBody traverses the given request body and all unique non-reference schemas in it.
func RequestBody(requestBody *openapi3.RequestBodyRef, visitor SchemaVisitor, levelNames ...string) error {
	if requestBody == nil {
		return nil
	}
	if requestBody.Ref != "" {
		return nil
	}
	if requestBody.Value == nil {
		return nil
	}
	if requestBody.Value.Content == nil {
		return nil
	}
	for contentType, mediaType := range requestBody.Value.Content {
		if mediaType == nil {
			continue
		}
		if err := MediaType(mediaType, visitor, append(levelNames, contentType, RequestSuffix)...); err != nil {
			return err
		}
	}
	return nil
}
