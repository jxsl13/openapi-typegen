package names_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/names"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	in := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	out := names.Deduplicate(in)

	assert.Equal(t, expected, out)

	in = []string{"a", "b", "c", "d", "ef", "f", "g", "h", "i", "j"}
	expected = []string{"a", "b", "c", "d", "ef", "g", "h", "i", "j"}
	out = names.Deduplicate(in)
	assert.Equal(t, expected, out)

	in = []string{"a", "bd", "c", "d", "ef", "f", "g", "gh", "i", "j"}
	expected = []string{"a", "bd", "c", "ef", "gh", "i", "j"}
	out = names.Deduplicate(in)
	assert.Equal(t, expected, out)

}
