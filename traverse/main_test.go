package traverse_test

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/require"
)

var Documents map[string]*openapi3.T

func TestMain(m *testing.M) {
	Documents = testutils.LoadSpecs(`\d{3,}.*\.yaml`, "../testdata/")

	rc := m.Run()
	os.Exit(rc)
}

func TestUniqueNames(t *testing.T) {
	names := make(map[string]bool, 128)

	for _, doc := range Documents {
		count := 0
		err := traverse.Parameters(doc, func(name string, parameter *openapi3.ParameterRef) error {
			require.NotNil(t, parameter)
			require.NotNil(t, parameter.Value)
			count++

			require.Falsef(t, names[name], "duplicate parameter name: %s", name)
			names[name] = true
			return nil
		})
		require.NoError(t, err)

		err = traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
			require.NotNil(t, schema)
			require.NotNil(t, schema.Value)
			count++

			require.Falsef(t, names[name], "duplicate schema name: %s", name)
			names[name] = true
			return nil
		})

		require.NoError(t, err)
	}

}

func TestTraverseParameters(t *testing.T) {
	var (
		err   error
		names = make(map[string]bool, 128)
	)

	for _, doc := range Documents {
		count := 0
		err = traverse.Parameters(doc, func(name string, parameter *openapi3.ParameterRef) error {
			require.NotNil(t, parameter)
			require.NotNil(t, parameter.Value)
			count++

			require.Falsef(t, names[name], "duplicate parameter name: %s", name)
			names[name] = true
			return nil
		})
		require.NoError(t, err)
	}

}

func TestTraverseSchemas(t *testing.T) {
	names := make(map[string]bool, 128)

	for _, doc := range Documents {

		count := 0
		err := traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
			require.NotNil(t, schema)
			require.NotNil(t, schema.Value)
			count++

			require.Falsef(t, names[name], "duplicate schema name: %s", name)
			names[name] = true
			return nil
		})
		require.NoError(t, err)
	}

}
