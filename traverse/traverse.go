package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var (
	RequestSuffix   = "Request"
	ResponseSuffix  = "Response"
	HeaderSuffix    = "Header"
	ParameterSuffix = "Parameter"
	SchemaSuffix    = "Schema"
)

type SchemaVisitor func(schemaRef *openapi3.SchemaRef, levelNames ...string) error

// Document traverses the given document and calls the visitor for each schema.
func Document(t *openapi3.T, visitor SchemaVisitor, levelNames ...string) error {
	if t == nil {
		return nil
	}

	var err error
	for pathName, pathItem := range t.Paths {
		if pathItem == nil {
			continue
		}
		if pathItem.Ref != "" {
			continue
		}

		err = PathItem(pathItem, visitor, append(levelNames, pathName)...)
		if err != nil {
			return err
		}
	}

	// traverse components
	if t.Components != nil {
		err = Components(t.Components, visitor, levelNames...)
		if err != nil {
			return err
		}
	}

	return nil
}
