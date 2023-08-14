package templates

import (
	"go/ast"
	"go/format"
	"os"
	"testing"

	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/stretchr/testify/require"
)

func TestStruct(t *testing.T) {

	fs, templates := testutils.LoadTemplates(`.*`, "../templates/")

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
