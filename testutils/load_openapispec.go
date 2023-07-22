package testutils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
)

func LoadSpec(relativeFilePath string) (doc *openapi3.T) {
	return loadSpec(FilePath(relativeFilePath))
}

func loadSpec(absolutePath string) (doc *openapi3.T) {
	loader := openapi3.NewLoader()

	doc, err := loader.LoadFromFile(absolutePath)
	if err != nil {
		panic(fmt.Errorf("failed to load openapi specification: %w", err))
	}

	err = doc.Validate(loader.Context)
	if err != nil {
		panic(fmt.Errorf("failed to validate openapi specification: %w", err))
	}
	return doc
}

func LoadSpecs(regex, dirPath string) map[string]*openapi3.T {
	dirPath = FilePath(dirPath)
	re := regexp.MustCompile(regex)

	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	result := make(map[string]*openapi3.T, 4)
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		fileName := fi.Name()
		if re.MatchString(fileName) {
			result[fileName] = loadSpec(filepath.Join(dirPath, fileName))
		}
	}
	return result
}
