package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// Schemas
func Schemas(doc *openapi3.T, visitor SchemaRefVisitor) error {
	var err error

	err = ComponentSchemas(doc, visitor)
	if err != nil {
		return err
	}

	err = ComponentParameters(doc, visitor)
	if err != nil {
		return err
	}

	err = ComponentHeaders(doc, visitor)
	if err != nil {
		return err
	}

	err = ComponentRequestBodies(doc, visitor)
	if err != nil {
		return err
	}

	err = ComponentResponses(doc, visitor)
	if err != nil {
		return err
	}

	err = PathHeaders(doc, visitor)
	if err != nil {
		return err
	}

	err = PathParameters(doc, visitor)
	if err != nil {
		return err
	}

	return nil
}

type SchemaRefVisitor func(name string, schema *openapi3.SchemaRef) error
