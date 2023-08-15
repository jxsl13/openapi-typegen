package identifiers

import "go/ast"

func New(name string) *ast.Ident {
	return ast.NewIdent(name)
}
