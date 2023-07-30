package main_test

import (
	"testing"

	main "github.com/jxsl13/openapi-typegen/cmd/openapi-typegen"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/stretchr/testify/require"
)

func TestHelp(t *testing.T) {
	out, err := testutils.Exec(main.NewRootCmd(), "-h")
	require.NoError(t, err)
	require.NotEmpty(t, out, "output is empty")
}

func Test001(t *testing.T) {
	out, err := testutils.Exec(main.NewRootCmd(), "-f", testutils.FilePath("../../testdata/001_schemas.yaml"))
	require.NoError(t, err)
	require.NotEmpty(t, out, "output is empty")
}
