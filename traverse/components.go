package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Components traverses the given components and all unique non-reference schemas in it.
func Components(components *openapi3.Components, visitor SchemaVisitor, levelNames ...string) error {
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
		err = visitor(schema, append(levelNames, schemaName, SchemaSuffix)...)
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

		err = Header(header, visitor, append(levelNames, headerName)...)
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
		err = Parameter(parameter, visitor, append(levelNames, parameterName)...)
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
		err = RequestBody(requestBody, visitor, append(levelNames, requestBodyName)...)
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
		err = Response(response, visitor, append(levelNames, responseName)...)
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
		err = Callback(callback, visitor, append(levelNames, callbackName)...)
		if err != nil {
			return err
		}
	}

	return nil
}
