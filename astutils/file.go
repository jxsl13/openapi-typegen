package astutils

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"io"
)

type File struct {
	name    string
	fs      *FileSet
	f       *ast.File
	imports []*ast.ImportSpec
}

func NewFile(name, packageName string) *File {
	return NewFileWithFileSet(NewFileSet(), name, packageName)
}

func NewFileWithFileSet(fs *FileSet, name, packageName string) *File {

	start := token.Pos(1)

	doc := NewCommentGroup(start, GeneratorComment())

	pkgName := &ast.Ident{
		NamePos: doc.End() + 1,
		Name:    packageName,
	}

	return &File{
		name: name,
		f: &ast.File{
			Doc:  doc,
			Name: pkgName,
		},
		fs: fs,
	}

}

func (f *File) Fprint(w io.Writer) error {

	err := printer.Fprint(w, f.fs.fs, f.f)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Print() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := f.Fprint(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f *File) PrintString() (string, error) {
	buf := &bytes.Buffer{}

	err := f.Fprint(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (f *File) String() string {
	s, err := f.PrintString()
	if err != nil {
		panic(err)
	}
	return s
}

func (f *File) Fformat(w io.Writer) error {
	return format.Node(w, f.fs.fs, f.f)
}

func (f *File) Format() ([]byte, error) {
	buf := &bytes.Buffer{}

	err := format.Node(buf, f.fs.fs, f.f)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f *File) Name() string {
	return f.name
}

func (f *File) AddImport(path string, alias ...string) {
	var a *ast.Ident
	if len(alias) > 0 && len(alias[0]) > 0 {
		a = ast.NewIdent(alias[0])
	}
	f.f.Imports = append(f.f.Imports, &ast.ImportSpec{
		Name: a,
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: path,
		},
	})

	ast.SortImports(nil, f.f)
}
