package traverse_test

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/assert"
)

func TestSingleParameterTypeMustHaveInKey(t *testing.T) {
	doc := Documents["004_callbacks.yaml"]

	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {

		if values, ok := levelNames[traverse.TypeKey]; ok && len(values) > 0 && values[0] == traverse.ParameterType {
			assert.Contains(t, levelNames, traverse.InKey)
		}
		return nil
	})
}

func TestAllParameterTypeMustHaveInKey(t *testing.T) {

	for _, doc := range OrderedDocuments {
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			if values, ok := levelNames[traverse.TypeKey]; ok && len(values) > 0 && values[0] == traverse.ParameterType {
				assert.Contains(t, levelNames, traverse.InKey)
			}
			return nil
		})
	}
}

func TestSingleParameterMustHaveExactlyOneInValue(t *testing.T) {
	doc := Documents["002_parameters.yaml"]

	cnt := 0
	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {

		if values, ok := levelNames[traverse.InKey]; ok {
			assert.Len(t, values, 1)
			cnt++
		}
		return nil
	})
	t.Logf("found %d parameters that have a in key", cnt)
}

func TestAllParameterMustHaveExactlyOneInValue(t *testing.T) {

	for _, doc := range OrderedDocuments {
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			if values, ok := levelNames[traverse.TypeKey]; ok {
				assert.Len(t, values, 1)
			}
			return nil
		})
	}
}
