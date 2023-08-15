package comment_test

import (
	"testing"

	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/tree/comment"
	"github.com/stretchr/testify/require"
)

func TestDoc(t *testing.T) {
	doc := comment.Doc("a", "b", "c")
	testutils.Dump(doc)

	text := doc.Text()
	require.NotEmpty(t, text)
}
