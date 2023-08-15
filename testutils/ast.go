package testutils

import (
	"go/ast"
	"go/token"
)

func Dump(node ast.Node) {
	err := ast.Print(token.NewFileSet(), node)
	if err != nil {
		panic(err)
	}
}
