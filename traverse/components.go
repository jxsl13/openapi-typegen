package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Components traverses the given components and all unique non-reference schemas in it.
func Components(components *openapi3.Components, visitor SchemaVisitor, levelNames map[string][]string) error {
	if components == nil {
		return nil
	}

	var err error

	// traverse schemas
	for schemaName, schema := range components.Schemas {
		if schema == nil {
			continue
		}
		if schema.Ref != "" {
			continue
		}
		err = visitor(schema, add(levelNames, NameKey, schemaName, TypeKey, SchemaType))
		if err != nil {
			return err
		}
	}

	// traverse headers
	for headerName, header := range components.Headers {
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

	// traverse parameters
	for parameterName, parameter := range components.Parameters {
		if parameter == nil {
			continue
		}
		if parameter.Ref != "" {
			continue
		}
		err = Parameter(parameter, visitor, add(levelNames, NameKey, parameterName))
		if err != nil {
			return err
		}
	}

	// traverse request bodies
	for requestBodyName, requestBody := range components.RequestBodies {
		if requestBody == nil {
			continue
		}
		if requestBody.Ref != "" {
			continue
		}
		err = RequestBody(requestBody, visitor, add(levelNames, NameKey, requestBodyName))
		if err != nil {
			return err
		}
	}

	// traverse responses
	for responseName, response := range components.Responses {
		if response == nil {
			continue
		}
		if response.Ref != "" {
			continue
		}
		err = Response(response, visitor, add(levelNames, NameKey, responseName))
		if err != nil {
			return err
		}
	}

	// traverse callbacks
	for callbackName, callback := range components.Callbacks {
		if callback == nil {
			continue
		}
		if callback.Ref != "" {
			continue
		}
		err = Callback(callback, visitor, add(levelNames, NameKey, callbackName))
		if err != nil {
			return err
		}
	}

	return nil
}
