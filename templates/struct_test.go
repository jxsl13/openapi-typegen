package templates_test

import (
	"embed"
	"go/ast"
	"go/format"
	"os"
	"testing"

	"github.com/jxsl13/openapi-typegen/fsutils"
	"github.com/jxsl13/openapi-typegen/templates"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/stretchr/testify/require"
)

func TestOsTemplates(t *testing.T) {
	fs, templates := testutils.LoadTemplates()

	require.NotEmpty(t, templates)

	for _, file := range templates {

		err := ast.Print(fs, file)
		if err != nil {
			t.Fatal(err)
		}

		// Print out parsed code to check round-tripping.
		err = format.Node(os.Stdout, fs, file)
		if err != nil {
			t.Fatal(err)
		}
	}
}

//go:embed *.go
var templateFs embed.FS

func TestFilePaths(t *testing.T) {
	root, names, err := fsutils.FilePaths(templateFs, ".*", ".")
	require.NoError(t, err)
	require.NotEmpty(t, root)
	require.NotEmpty(t, names)
}

func TestEmbeddedTemplates(t *testing.T) {
	fs, templates := templates.FileSet, templates.Templates

	require.NotEmpty(t, templates)

	for _, file := range templates {

		err := ast.Print(fs, file)
		if err != nil {
			t.Fatal(err)
		}

		// Print out parsed code to check round-tripping.
		err = format.Node(os.Stdout, fs, file)
		if err != nil {
			t.Fatal(err)
		}
	}
}
