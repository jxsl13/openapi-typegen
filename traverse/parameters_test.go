package traverse_test

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/require"
)

func TestTraverseParameters(t *testing.T) {
	doc := Documents["003_components.yaml"]

	count := 0
	err := traverse.Parameters(doc, func(name string, parameter *openapi3.ParameterRef) error {
		require.NotNil(t, parameter)
		require.NotNil(t, parameter.Value)
		count++
		return nil
	})
	require.NoError(t, err)
}

func TestTraverseSchemas(t *testing.T) {
	doc := Documents["001_mangadex.yaml"]

	count := 0
	err := traverse.Schemas(doc, func(name string, schema *openapi3.SchemaRef) error {
		require.NotNil(t, schema)
		require.NotNil(t, schema.Value)
		count++
		return nil
	})
	require.NoError(t, err)
}
