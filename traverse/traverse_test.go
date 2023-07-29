package traverse_test

import (
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/testutils"
	"github.com/jxsl13/openapi-typegen/traverse"
	"github.com/stretchr/testify/require"
)

var (
	Documents        = testutils.LoadSpecs(`\d+.*.yaml`, "../testdata/")
	OrderedDocuments = mapToOrderedTupleList(Documents)

	SuffixMap = map[string]bool{
		traverse.RequestSuffix:   true,
		traverse.ResponseSuffix:  true,
		traverse.HeaderSuffix:    true,
		traverse.ParameterSuffix: true,
		traverse.SchemaSuffix:    true,
	}
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

func TestTraverse(t *testing.T) {
	for _, tuple := range OrderedDocuments {
		cnt := 0
		//unique := make(map[string]bool, 64)
		lenBuckets := make(map[int]int, 64)

		err := traverse.Document(tuple.Doc, func(schemaRef *openapi3.SchemaRef, levelNames ...string) error {
			t.Logf("document: %s, levelNames: %v", tuple.Name, levelNames)
			require.Greater(t, len(levelNames), 0)
			cnt++

			lenBuckets[len(levelNames)]++
			return nil
		})
		t.Logf("document: %s, cnt: %d", tuple.Name, cnt)
		t.Logf("document: %s, lenBuckets: %v", tuple.Name, lenBuckets)
		require.NoError(t, err)

	}
}
