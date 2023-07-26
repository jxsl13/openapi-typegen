package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func ComponentSchemas(doc *openapi3.T, visitor SchemaRefVisitor) error {
	var err error
	if doc.Components == nil {
		return nil
	}
	for name, schema := range doc.Components.Schemas {

		if schema.Ref != "" || schema.Value == nil {
			continue
		}
		// only traverse on-references
		err = visitor(names.ToTitleTypeName(name), schema)
		if err != nil {
			return err
		}
	}
	return nil
}
