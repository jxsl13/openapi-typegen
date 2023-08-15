package identifiers_test

import (
	"fmt"
	"testing"

	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/tree/identifiers"
	"github.com/stretchr/testify/require"
)

func TestIdentifiers(t *testing.T) {
	var (
		in       = "somName"
		expected = fmt.Sprint(in)
	)
	out := testutils.NodeString(identifiers.New(in))
	require.Equal(t, expected, out)

}
