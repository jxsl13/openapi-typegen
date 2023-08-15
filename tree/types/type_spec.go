package types

import (
	"go/ast"
	"go/token"

	"github.com/jxsl13/openapi-typegen/tree/comment"
)

// TODO: WIP - design might change, as this does not seem to work as expected
func NewTypeSpec(name, typeName string, comments ...string) *ast.TypeSpec {

	doc := comment.Doc(comments...)

	nameIdent := &ast.Ident{
		NamePos: doc.End() + token.Pos(len("import")+1), // allow inserting "type", "import" and "const" keywords
		Name:    name,
	}

	typeIfdent := &ast.Ident{
		NamePos: nameIdent.End() + 1,
		Name:    typeName,
	}

	spec := &ast.TypeSpec{
		Doc:  doc,
		Name: nameIdent,
		Type: typeIfdent,
	}

	return spec
}
