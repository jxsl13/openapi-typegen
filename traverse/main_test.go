package traverse_test

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/testutils"
)

var Documents map[string]*openapi3.T

func TestMain(m *testing.M) {
	Documents = testutils.LoadSpecs(`\d{3,}.*\.yaml`, "../testdata/")

	rc := m.Run()
	os.Exit(rc)

}
