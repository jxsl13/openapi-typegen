package traverse

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func Schemas(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil || doc.Components.Schemas == nil {
		return nil
	}

	var err error
	schemas := doc.Components.Schemas

	for k, v := range schemas {
		if v.Value == nil {
			panic(fmt.Sprintf("component schema %q is nil", k))
		}
		err = visitor(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func schemaByRefName(schemas openapi3.Schemas, ref string) (*openapi3.Schema, error) {
	sr, ok := schemas[ref]
	if !ok {
		return nil, fmt.Errorf("could not find reference %q in component schemas", ref)
	} else if sr.Value == nil {
		return nil, fmt.Errorf("could reference %q in component schemas is nil", ref)
	}

	return sr.Value, nil
}

type SchemaRefVisitor func(name string, schema *openapi3.SchemaRef) error
