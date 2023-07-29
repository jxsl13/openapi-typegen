package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Responses traverses the given responses and all unique non-reference schemas in it.
func Responses(responses openapi3.Responses, visitor SchemaVisitor, levelNames ...string) error {
	for statusCode, response := range responses {
		if response == nil {
			continue
		}
		if err := Response(response, visitor, append(levelNames, statusCode)...); err != nil {
			return err
		}
	}
	return nil
}

// Response traverses the given response and all unique non-reference schemas in it.
func Response(response *openapi3.ResponseRef, visitor SchemaVisitor, levelNames ...string) error {
	if response == nil {
		return nil
	}
	if response.Ref != "" {
		return nil
	}
	if response.Value == nil {
		return nil
	}

	//traverse headers
	var err error
	for headerName, header := range response.Value.Headers {
		if header == nil {
			continue
		}
		if header.Ref != "" {
			continue
		}
		err = Header(header, visitor, append(levelNames, headerName)...)
		if err != nil {
			return err
		}
	}

	if response.Value.Content == nil {
		return nil
	}
	for contentType, mediaType := range response.Value.Content {
		if mediaType == nil {
			continue
		}
		if err := MediaType(mediaType, visitor, append(levelNames, contentType, ResponseSuffix)...); err != nil {
			return err
		}
	}
	return nil
}
