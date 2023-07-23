package traverse_test

import (
	"os"
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

func TestUniqueNames(t *testing.T) {

	for k, doc := range Documents {
		t.Logf("document: %s", k)
		names := make(map[string]bool, 128)
		count := 0

		t.Logf("parameters: %s", k)
		err := traverse.Parameters(doc, func(name string, parameter *openapi3.ParameterRef) error {
			require.NotRegexp(t, `^\d+`, name, "name starts with integer")
			require.NotNil(t, parameter)
			require.NotNil(t, parameter.Value)
			require.Empty(t, parameter.Ref)
			count++

			assert.Falsef(t, names[name], "duplicate parameter name: %s", name)
			names[name] = true
			t.Logf("parameter name = %s", name)
			return nil
		})
		require.NoError(t, err)

		t.Logf("schemas: %s", k)
		err = traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
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
	}

}

func TestTraverseParameters(t *testing.T) {
	var err error

	for k, doc := range Documents {
		t.Logf("document: %s", k)
		names := make(map[string]bool, 128)
		count := 0
		err = traverse.Parameters(doc, func(name string, parameter *openapi3.ParameterRef) error {
			require.NotRegexp(t, `^\d+`, name, "name starts with integer")
			require.NotNil(t, parameter)
			require.NotNil(t, parameter.Value)
			require.Empty(t, parameter.Ref)
			count++

			assert.Falsef(t, names[name], "duplicate parameter name: %s", name)
			names[name] = true
			t.Logf("parameter name = %s", name)
			return nil
		})
		require.NoError(t, err)
	}

}

func TestTraverseSchemas(t *testing.T) {

	for k, doc := range Documents {
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
	}

}
