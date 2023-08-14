package literals_test

import (
	"bytes"
	"fmt"
	"go/printer"
	"go/token"
	"testing"

	"github.com/jxsl13/openapi-typegen/tree/literals"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {

	in := 10
	out := fmt.Sprint(in)

	lit := literals.Int(in)
	buf := &bytes.Buffer{}
	printer.Fprint(buf, token.NewFileSet(), lit)
	require.Equal(t, out, buf.String())
}

func TestFloat(t *testing.T) {

	in := 10.5
	out := fmt.Sprint(in)

	lit := literals.Float(in)
	buf := &bytes.Buffer{}
	printer.Fprint(buf, token.NewFileSet(), lit)
	require.Equal(t, out, buf.String())
}

func TestString(t *testing.T) {

	in := "hello world"
	out := fmt.Sprintf("%q", in)

	lit := literals.String(in)
	buf := &bytes.Buffer{}
	printer.Fprint(buf, token.NewFileSet(), lit)
	require.Equal(t, out, buf.String())
}
