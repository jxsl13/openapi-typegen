package traverse

import "github.com/getkin/kin-openapi/openapi3"

// Callback traverses all unique non-reference schemas in the given callback.
func Callback(callback *openapi3.CallbackRef, visitor SchemaVisitor, levelNames ...string) error {
	if callback == nil {
		return nil
	}
	if callback.Ref != "" {
		return nil
	}
	if callback.Value == nil {
		return nil
	}

	var err error
	for callbackName, pathItem := range *callback.Value {
		if pathItem == nil {
			continue
		}
		if pathItem.Ref != "" {
			continue
		}

		err = PathItem(pathItem, visitor, append(levelNames, callbackName)...)
		if err != nil {
			return err
		}
	}

	return nil
}
