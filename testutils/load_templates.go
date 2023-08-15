package testutils

import (
	"go/ast"
	"go/token"

	"github.com/jxsl13/openapi-typegen/fsutils"
	"github.com/jxsl13/openapi-typegen/templates"
)

// LoadTemplates loads all .go files starting with template_ or ending with _template.go
// from the templates folder
func LoadTemplates() (*token.FileSet, map[string]*ast.File) {

	const dirPath = "../templates/"
	fs, files, err := templates.LoadAll(fsutils.NewOsFS(), templates.NameRegex, FilePath(dirPath))
	if err != nil {
		panic(err)
	}
	return fs, files
}
