package comment

import (
	"fmt"
	"go/ast"
	"go/token"
)

func Doc(lines ...string) *ast.CommentGroup {
	var cg *ast.CommentGroup = &ast.CommentGroup{}

	pos := token.Pos(1) // 0 is NoPos
	for _, line := range lines {
		if len(cg.List) > 0 {
			pos = cg.End()
		}

		// line directives
		cg.List = append(cg.List, &ast.Comment{
			Slash: pos,
			Text:  fmt.Sprintf("//line :%d", pos),
		})

		cg.List = append(cg.List, &ast.Comment{
			Slash: pos,
			Text:  "// " + line,
		})
	}

	return cg
}
