package spec

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/fsutils"
)

var FileRegex = ".(yml|yaml|json)$"

// Paths returns the absolute dir path and a list of template file names
func Paths(regex, dirPath string) (string, []string, error) {

	re, err := regexp.Compile(regex)
	if err != nil {
		return "", nil, fmt.Errorf("failed to compile regex %q: %w", regex, err)
	}

	dir, files, err := fsutils.FilePaths(fsutils.NewOsFS(), regex, dirPath)
	if err != nil {
		return "", nil, err
	}
	result := make([]string, 0, len(files))
	for _, file := range files {
		if re.MatchString(file) {
			result = append(result, file)
		}
	}
	return dir, result, nil
}

// Load loads a single spec from the given relative or absolute path
func Load(relative string, options ...Option) (doc *openapi3.T, err error) {
	opt := loadOptions{
		validationOptions: []openapi3.ValidationOption{
			openapi3.EnableSchemaDefaultsValidation(),
			openapi3.EnableSchemaPatternValidation(),
			openapi3.DisableSchemaFormatValidation(), // might at some point introduce custom format types
		},
	}

	absolutePath, err := fsutils.Absolute(relative)
	if err != nil {
		return nil, err
	}

	loader := openapi3.Loader{
		IsExternalRefsAllowed: true,
		Context:               context.Background(),
	}

	doc, err = loader.LoadFromFile(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load openapi specification %s: %w", absolutePath, err)
	}

	err = doc.Validate(
		loader.Context,
		opt.validationOptions...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to validate openapi specification %s: %w", absolutePath, err)
	}
	return doc, nil
}

// LoadAll loads all specs that match the given regex from the given dirPath
func LoadAll(regex, dirPath string, options ...Option) (map[string]*openapi3.T, error) {

	dirPath, matchingFiles, err := Paths(regex, dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file paths: %w", err)
	}

	result := make(map[string]*openapi3.T, len(matchingFiles))
	for _, filename := range matchingFiles {
		spec, err := Load(filepath.Join(dirPath, filename), options...)
		if err != nil {
			return nil, err
		}
		result[filename] = spec
	}
	return result, nil
}
