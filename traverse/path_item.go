package traverse

import "github.com/getkin/kin-openapi/openapi3"

// PathItem traverses the given path item and calls the visitor for each schema.
func PathItem(pathItem *openapi3.PathItem, visitor SchemaVisitor, levelNames map[string][]string) error {
	if pathItem == nil {
		return nil
	}
	if pathItem.Ref != "" {
		return nil
	}

	// travserse path item parameters
	var err error
	for _, parameter := range pathItem.Parameters {
		if parameter == nil {
			continue
		}
		if parameter.Ref != "" {
			continue
		}
		err = Parameter(parameter, visitor, levelNames)
		if err != nil {
			return err
		}
	}

	for method, operation := range pathItem.Operations() {
		if operation == nil {
			continue
		}

		err = Operation(operation, visitor, add(levelNames, MethodKey, method, OperationKey, operation.OperationID))
		if err != nil {
			return err
		}
	}

	return nil
}
