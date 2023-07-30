package traverse_test

import (
	"os"
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/options"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/k0kubun/pp/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	Documents        = testutils.LoadSpecs(`\d+.*.yaml`, "../testdata/")
	OrderedDocuments = mapToOrderedTupleList(Documents)
)

type Tuple struct {
	Name string
	Doc  *openapi3.T
}

type byTupleName []Tuple

func (a byTupleName) Len() int           { return len(a) }
func (a byTupleName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byTupleName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func mapToOrderedTupleList(m map[string]*openapi3.T) []Tuple {
	tuples := make([]Tuple, 0, len(m))
	for name, doc := range m {
		tuples = append(tuples, Tuple{
			Name: name,
			Doc:  doc,
		})
	}

	sort.Sort(byTupleName(tuples))
	return tuples
}

func TestMain(m *testing.M) {
	pp.Default.SetColoringEnabled(false)
	os.Exit(m.Run())
}

func TestSingleMustContainTypeKey(t *testing.T) {
	doc := Documents["004_callbacks.yaml"]

	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
		assert.Contains(t, levelNames, traverse.TypeKey)
		return nil
	})
}

func TestAllMustContainTypeKey(t *testing.T) {

	for _, doc := range OrderedDocuments {
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			assert.Contains(t, levelNames, traverse.TypeKey)
			return nil
		})
	}
}

func TestSingleMustContainOneTypeKey(t *testing.T) {
	doc := Documents["004_callbacks.yaml"]

	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
		//must only contain one type name
		assert.Equal(t, len(levelNames[traverse.TypeKey]), 1)
		return nil
	})
}

func TestAllMustContainOneTypeKey(t *testing.T) {

	for _, doc := range OrderedDocuments {
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			//must only contain one type name
			assert.Equal(t, len(levelNames[traverse.TypeKey]), 1)
			return nil
		})
	}
}

func TestSingleUniqueOperation(t *testing.T) {
	doc := Documents["002_parameters.yaml"]

	namesMap := make(map[string]bool)
	traverse.Document(doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
		name := options.UniqueName(levelNames)
		t.Logf("unique name: %s", name)
		require.NotEmpty(t, name, "name is empty")
		require.NotContains(t, namesMap, name)
		namesMap[name] = true

		return nil
	})
}

func TestAllUniqueOperation(t *testing.T) {

	for _, doc := range OrderedDocuments {
		namesMap := make(map[string]bool)
		stackTraceMap := make(map[string]string)
		traverse.Document(doc.Doc, func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error {
			name := options.UniqueName(levelNames)
			require.NotEmpty(t, name, "name is empty")
			t.Logf("document: %s unique name: %s", doc.Name, name)
			require.NotContainsf(t, namesMap, name, "previously seen at: %s", stackTraceMap[name])
			namesMap[name] = true
			stackTraceMap[name] = testutils.StackTraceString()

			return nil
		})
	}

}
