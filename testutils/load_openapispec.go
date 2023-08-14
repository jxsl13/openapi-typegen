package testutils

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/spec"
)

// LoadSpec loads an openapi spec relative to the current go file's location.
func LoadSpec(relativeFilePath string) (doc *openapi3.T) {
	doc, err := spec.Load(FilePath(relativeFilePath))
	if err != nil {
		panic(err)
	}
	return doc
}

// LoadSpecs looks into the directory path relative to the current go file's directory path.
// You may pass a regular expression to match a specific subset of files in the passed directory.
func LoadSpecs(regex, dirPath string) map[string]*openapi3.T {
	specs, err := spec.LoadAll(regex, FilePath(dirPath))
	if err != nil {
		panic(err)
	}
	return specs
}
