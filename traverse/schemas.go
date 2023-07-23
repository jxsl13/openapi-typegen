package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

// Schemas traverses #/components/schemas
func Schemas(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil {
		return nil
	}

	var err error
	schemas := doc.Components.Schemas

	for k, v := range schemas {
		if v.Ref != "" {
			continue
		}
		if v.Value == nil {
			continue
		}

		name := names.ToTitleTypeName(k)
		err = visitor(name, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type SchemaRefVisitor func(name string, schema *openapi3.SchemaRef) error
