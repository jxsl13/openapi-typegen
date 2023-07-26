package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func ComponentHeaders(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil {
		return nil
	}
	var err error
	for name, header := range doc.Components.Headers {
		if header.Ref != "" || header.Value == nil {
			continue
		}

		if header.Value.Schema == nil ||
			header.Value.Schema.Ref != "" ||
			header.Value.Schema.Value == nil {
			continue
		}

		err = visitor(names.ToTitleTypeName(name), header.Value.Schema)
		if err != nil {
			return err
		}
	}

	return nil
}
