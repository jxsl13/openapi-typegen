package names_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/names"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicateEqual(t *testing.T) {
	var (
		in       = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		expected = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		out      = names.Deduplicate(in)
	)
	assert.Equal(t, expected, out)
}

func TestDeduplicateOne(t *testing.T) {

	var (
		in       = []string{"a", "b", "c", "d", "ef", "F", "g", "h", "i", "j"}
		expected = []string{"a", "b", "c", "d", "ef", "g", "h", "i", "j"}
		out      = names.Deduplicate(in)
	)
	assert.Equal(t, expected, out)
}

func TestDeduplicateThree(t *testing.T) {

	var (
		in       = []string{"a", "bd", "c", "d", "ef", "F", "g", "Gh", "i", "j"}
		expected = []string{"a", "bd", "c", "ef", "Gh", "i", "j"}
		out      = names.Deduplicate(in)
	)
	assert.Equal(t, expected, out)
}

func TestDeDuplicateOneEqual(t *testing.T) {

	var (
		in       = []string{"A", "a"}
		expected = []string{"A"}
		out      = names.Deduplicate(in)
	)
	assert.Equal(t, expected, out)

	in = []string{"a", "A"}
	expected = []string{"a"}
	out = names.Deduplicate(in)

	assert.Equal(t, expected, out)
}
