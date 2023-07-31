package astutils_test

import (
	"go/token"
	"testing"

	"github.com/jxsl13/openapi-typegen/astutils"
	"github.com/stretchr/testify/require"
)

func TestNewComment(t *testing.T) {
	cg := astutils.NewCommentGroup(token.Pos(1), "test", "test2")
	require.Len(t, cg.List, 2)
	require.Equal(t, "test", cg.List[0].Text)
	require.Equal(t, "test2", cg.List[1].Text)
}
