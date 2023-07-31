package astutils

import (
	"go/ast"
	"go/token"
)

func NewCommentGroup(startPos token.Pos, lines ...string) *ast.CommentGroup {
	cg := &ast.CommentGroup{
		List: make([]*ast.Comment, 0, len(lines)),
	}

	pos := startPos
	for _, line := range lines {
		cg.List = append(cg.List, &ast.Comment{
			Slash: pos,
			Text:  line,
		})
		pos = cg.End() + 1
	}

	return cg
}
