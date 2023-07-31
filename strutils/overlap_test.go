package strutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOverlap(t *testing.T) {
	var i int
	i = Overlap("abcdefghijklmnop", "ghijklmnopqrstuvw")
	require.Equal(t, 10, i)

	i = Overlap("abcdefg", "hijklabcdg")
	require.Equal(t, 0, i)

	i = Overlap("abcdefghijklmnopabcdefghijklmnop", "ghijklmnopqrstuvwxyzabcdefghijklmnop")
	require.Equal(t, 10, i)
}

func FuzzOverlap(f *testing.F) {
	f.Add("abcdefghijklmnop", "ghijklmnopqrstuvw")
	f.Add("abcdefg", "hijklabcdg")
	f.Add("abcdefghijklmnopabcdefghijklmnop", "ghijklmnopqrstuvwxyzabcdefghijklmnop")

	f.Fuzz(func(t *testing.T, a, b string) {
		i := Overlap(a, b)
		if i == 0 {
			return
		}

		require.Equal(t, a[len(a)-i:], b[:i])
	})
}
