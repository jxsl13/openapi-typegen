package strutils_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/strutils"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	s := strutils.Merge("abcABC", "ABCdefg")
	require.Equal(t, "abcABCdefg", s)
}

func TestMergeAll(t *testing.T) {
	s := strutils.MergeAll("abcABC", "ABCdefg", "fghijk")
	require.Equal(t, "abcABCdefghijk", s)
}
