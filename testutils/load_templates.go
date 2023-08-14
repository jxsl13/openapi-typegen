package testutils

import (
	"go/ast"
	"go/token"

	"github.com/jxsl13/openapi-typegen/template"
)

// LoadTemplates loads all .go files that are not test files from the given relative directory path which is relative to the current source file directory.
func LoadTemplates(regex, dirPath string) (*token.FileSet, map[string]*ast.File) {
	fs, files, err := template.Templates(regex, dirPath)
	if err != nil {
		panic(err)
	}
	return fs, files
}
