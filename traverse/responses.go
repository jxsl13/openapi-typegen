package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Responses traverses the given responses and all unique non-reference schemas in it.
func Responses(responses openapi3.Responses, visitor SchemaVisitor, levelNames map[string][]string) error {
	var err error
	for statusCode, response := range responses {
		if response == nil {
			continue
		}
		err = Response(response, visitor, add(levelNames, StatusKey, statusCode))
		if err != nil {
			return err
		}
	}
	return nil
}

// Response traverses the given response and all unique non-reference schemas in it.
func Response(response *openapi3.ResponseRef, visitor SchemaVisitor, levelNames map[string][]string) error {
	if response == nil {
		return nil
	}
	if response.Ref != "" {
		return nil
	}

	return ResponseSchema(response.Value, visitor, levelNames)
}

func ResponseSchema(response *openapi3.Response, visitor SchemaVisitor, levelNames map[string][]string) error {
	if response == nil {
		return nil
	}

	var err error
	for headerName, header := range response.Headers {
		if header == nil {
			continue
		}
		if header.Ref != "" {
			continue
		}
		err = Header(header, visitor, add(levelNames, NameKey, headerName))
		if err != nil {
			return err
		}
	}

	for contentType, mediaType := range response.Content {
		if mediaType == nil {
			continue
		}
		err = MediaType(mediaType, visitor, add(levelNames, ContentKey, contentType, TypeKey, ResponseType))
		if err != nil {
			return err
		}
	}
	return nil
}
