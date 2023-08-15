package testutils

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

func NodeString(n ast.Node) string {
	return NodeWithFileSetString(n, token.NewFileSet())
}

func NodeWithFileSetString(n ast.Node, fset *token.FileSet) string {
	buf := &bytes.Buffer{}
	printer.Fprint(buf, fset, n)
	return buf.String()
}
