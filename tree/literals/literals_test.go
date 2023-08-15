package literals_test

import (
	"fmt"
	"testing"

	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/tree/literals"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {

	var (
		in       = 10
		expected = fmt.Sprint(in)
	)

	require.Equal(t, expected, testutils.NodeString(literals.Int(in)))
}

func TestFloat(t *testing.T) {

	var (
		in       = 10.5
		expected = fmt.Sprint(in)
	)

	require.Equal(t, expected, testutils.NodeString(literals.Float(in)))
}

func TestString(t *testing.T) {

	var (
		in       = "hello world"
		expected = fmt.Sprintf("%q", in)
	)

	require.Equal(t, expected, testutils.NodeString(literals.String(in)))
}
