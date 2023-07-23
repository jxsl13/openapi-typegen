package names_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/names"
	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	s := names.Join("abcABC", "ABCdefg")
	require.Equal(t, "abcABCdefg", s)
}
