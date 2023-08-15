package templates

import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jxsl13/openapi-typegen/fsutils"
)

//go:embed *.go
var templateFs embed.FS

var (
	FileSet   *token.FileSet
	Templates map[string]*ast.File
	NameRegex = `(^template_|_template.go$)`
)

func init() {
	fs, tpl, err := LoadAll(templateFs, NameRegex, ".")
	if err != nil {
		panic(fmt.Errorf("failed to load embedded templates: %w", err))
	}

	FileSet = fs
	Templates = tpl
}

// Paths returns the absolute dir path and a list of template file names.
// Template files are .go files that match the regex and do not end with _test.go
func Paths(fsys fs.FS, regex, dirPath string) (string, []string, error) {

	re, err := regexp.Compile(regex)
	if err != nil {
		return "", nil, fmt.Errorf("failed to compile regex %q: %w", regex, err)
	}

	dir, files, err := fsutils.FilePaths(fsys, ".go$", dirPath)
	if err != nil {
		return "", nil, err
	}

	goFiles := make([]string, 0, len(files))
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		} else if re.MatchString(file) {
			goFiles = append(goFiles, file)
		}
	}

	return dir, goFiles, nil
}

// LoadAll returns a map of file names to ast.File pointers.
// and a token.FileSet which is needed for comment positions.
func LoadAll(fsys fs.FS, regex, dirPath string) (*token.FileSet, map[string]*ast.File, error) {
	dir, files, err := Paths(fsys, regex, dirPath)
	if err != nil {
		return nil, nil, err
	}

	//bits := parser.ParseComments | parser.SkipObjectResolution
	bits := parser.ParseComments

	fs := token.NewFileSet()

	result := make(map[string]*ast.File, len(files))
	for _, file := range files {
		data, err := fsutils.ReadFile(fsys, filepath.Join(dir, file))
		if err != nil {
			return nil, nil, err
		}
		f, err := parser.ParseFile(fs, file, data, bits)
		if err != nil {
			return nil, nil, err
		}
		result[file] = f
	}
	return fs, result, nil
}
