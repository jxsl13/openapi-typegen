package main_test

import (
	"testing"

	main "github.com/jxsl13/openapi-typegen/cmd/openapi-typegen"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/stretchr/testify/require"
)

func TestExec(t *testing.T) {
	out, err := testutils.Exec(main.NewRootCmd(), "-f", testutils.FilePath("../../testdata/001_petstore-expanded.yaml"))
	require.NoError(t, err)
	require.NotEmpty(t, out, "output is empty")
}
