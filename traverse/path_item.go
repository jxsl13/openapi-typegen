package traverse

import "github.com/getkin/kin-openapi/openapi3"

// PathItem traverses the given path item and calls the visitor for each schema.
func PathItem(pathItem *openapi3.PathItem, visitor SchemaVisitor, levelNames ...string) error {
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
		err = Parameter(parameter, visitor, levelNames...)
		if err != nil {
			return err
		}
	}

	for method, operation := range pathItem.Operations() {
		if operation == nil {
			continue
		}

		if err := Operation(operation, visitor, append(levelNames, method, operation.OperationID)...); err != nil {
			return err
		}
	}

	return nil
}
