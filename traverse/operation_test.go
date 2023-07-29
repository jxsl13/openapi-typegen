package traverse_test

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/assert"
)

func TestSingleOperationMustHavePath(t *testing.T) {
	doc := Documents["002_parameters.yaml"]

	cnt := 0
	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
		if values, ok := levelNames[traverse.OperationKey]; ok && len(values) > 0 && values[0] != "" {
			assert.Contains(t, levelNames, traverse.PathKey)
			assert.NotEmpty(t, levelNames[traverse.PathKey])
			cnt++
		}
		return nil
	})
	t.Logf("found %d operations that have a path key", cnt)
}

func TestAllOperationMustHavePath(t *testing.T) {

	cnt := 0
	for _, doc := range OrderedDocuments {
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			if values, ok := levelNames[traverse.OperationKey]; ok && len(values) > 0 && values[0] != "" {
				assert.Contains(t, levelNames, traverse.PathKey)
				assert.NotEmpty(t, levelNames[traverse.PathKey])
				cnt++
			}
			return nil
		})
	}
	t.Logf("found %d operations that have a path key", cnt)
}
