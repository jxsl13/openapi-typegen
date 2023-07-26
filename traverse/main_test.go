package traverse_test

import (
	"os"
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var Documents map[string]*openapi3.T

func TestMain(m *testing.M) {
	Documents = testutils.LoadSpecs(`\d{3,}.*\.yaml`, "../testdata/")

	rc := m.Run()
	os.Exit(rc)
}

func TestTraverseSchemas(t *testing.T) {

	expectedSchemas := map[string]int{
		"001_mangadex.yaml":   97,
		"002_parameters.yaml": 2,
		"003_components.yaml": 43,
		"004_schemas.yaml":    9,
	}

	iterateDocuments(func(k string, doc *openapi3.T) error {
		t.Logf("document: %s", k)
		names := make(map[string]bool, 128)

		count := 0
		err := traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
			require.NotRegexp(t, `^\d+`, name, "name starts with integer")
			require.NotNil(t, schema)
			require.NotNil(t, schema.Value)
			require.Empty(t, schema.Ref)
			count++

			assert.Falsef(t, names[name], "duplicate schema name: %s", name)
			names[name] = true
			t.Logf("schema name = %s", name)
			return nil
		})
		require.NoError(t, err)
		t.Logf("total schemas: %d", count)

		require.Equal(t, expectedSchemas[k], count)

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
