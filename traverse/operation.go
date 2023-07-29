package traverse

import "github.com/getkin/kin-openapi/openapi3"

func Operation(operation *openapi3.Operation, visitor SchemaVisitor, levelNames ...string) error {
	if operation == nil {
		return nil
	}

	// traverse parameters
	var err error
	for _, parameter := range operation.Parameters {
		if parameter == nil {
			continue
		}
		if parameter.Ref != "" {
			continue
		}
		err = Parameter(parameter, visitor, levelNames...)
		if err != nil {
			return err
		}
	}

	if operation.RequestBody != nil && operation.RequestBody.Ref == "" {
		if err := RequestBody(operation.RequestBody, visitor, append(levelNames, RequestSuffix)...); err != nil {
			return err
		}
	}
	if operation.Responses != nil {
		if err := Responses(operation.Responses, visitor, levelNames...); err != nil {
			return err
		}
	}

	// traverse callbacks
	for callbackName, callback := range operation.Callbacks {
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
