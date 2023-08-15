package types

import (
	"go/ast"
	"go/token"

	"github.com/jxsl13/openapi-typegen/tree/comment"
)

const (
	Bool       Primitive = "bool"
	Float32    Primitive = "float32"
	Float64    Primitive = "float64"
	String     Primitive = "string"
	Byte       Primitive = "byte"
	Int        Primitive = "int"
	Int8       Primitive = "int8"
	Int16      Primitive = "int16"
	Int32      Primitive = "int32"
	Int64      Primitive = "int64"
	Rune       Primitive = "rune"
	Uint16     Primitive = "uint16"
	Uint32     Primitive = "uint32"
	Uint64     Primitive = "uint64"
	Uint       Primitive = "uint"
	Uintptr    Primitive = "uintptr"
	Complex64  Primitive = "complex64"
	Complex128 Primitive = "complex128"
)

type Primitive string

// TODO: WIP - design might change, as this does not seem to work as expected
func NewPrimitive(name string, t Primitive, comments ...string) *ast.GenDecl {

	//size := len("type") + 1 + len(name) + 1 + len(string(t)) + strutils.Size(comments...) + len(comments)*4 // '// ' + '\n'
	cg := comment.Doc(comments...)

	spec := &ast.TypeSpec{
		Doc: cg,
	}

	spec.Type = &ast.Ident{
		NamePos: spec.End() + 1,
		Name:    string(t),
	}

	spec.Name = &ast.Ident{
		NamePos: spec.End() + 1,
		Name:    name,
	}

	primitive := &ast.GenDecl{
		TokPos: 1,
		Tok:    token.TYPE,
		Specs: []ast.Spec{
			spec,
		},
	}

	return primitive
}

func NewInt(name string, comments ...string) *ast.GenDecl {
	return NewPrimitive(name, Int, comments...)
}
