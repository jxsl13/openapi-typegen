package astutils_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/astutils"
	"github.com/stretchr/testify/require"
)

func TestGeneratorName(t *testing.T) {
	comment := astutils.GeneratorComment()
	require.NotEmpty(t, comment)
	require.NotContains(t, comment, "unknown generator")
}

func TestFile(t *testing.T) {

	file := astutils.NewFile("test.go", "test")
	file.AddImport("github.com/jxsl13/openapi-typegen/astutils")
	file.AddImport("fmt")

	code := file.String()
	require.NotEmpty(t, code)
}
