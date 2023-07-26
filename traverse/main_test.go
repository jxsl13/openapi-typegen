package traverse_test

import (
	"os"
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/assert"
)

var Documents map[string]*openapi3.T

func TestMain(m *testing.M) {
	Documents = testutils.LoadSpecs(`\d{3,}.*\.yaml`, "../testdata/")

	rc := m.Run()
	os.Exit(rc)
}

func TestTraverseSchemas(t *testing.T) {

	expectedSchemas := map[string]int{
		"001_schemas.yaml":    9,
		"002_parameters.yaml": 2,
		"003_components.yaml": 43,
		"100_mangadex.yaml":   97,
	}

	iterateDocuments(func(k string, doc *openapi3.T) error {
		t.Logf("document: %s", k)
		names := make(map[string]bool, 128)

		count := 0
		err := traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
			assert.NotRegexp(t, `^\d+`, name, "name starts with integer")
			assert.NotNil(t, schema)
			assert.NotNil(t, schema.Value)
			assert.Empty(t, schema.Ref)
			count++

			assert.Falsef(t, names[name], "duplicate schema name: %s", name)
			names[name] = true

			v := schema.Value
			t.Logf("schema name = %s (type=%s format=%s)", name, v.Type, v.Format)
			return nil
		})
		assert.NoError(t, err)
		t.Logf("total schemas: %d", count)

		assert.Equal(t, expectedSchemas[k], count)

		return nil
	})
}

func iterateSorted[V any](m map[string]V, visitor func(key string, value V) error) error {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var err error
	for _, k := range keys {
		err = visitor(k, m[k])
		if err != nil {
			return err
		}
	}
	return nil
}

func iterateDocuments(visitor func(name string, doc *openapi3.T) error) error {
	return iterateSorted(Documents, visitor)
}
