package filepath

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Absolute returns the absolute path of the given relative path.
func Absolute(relative string) (string, error) {
	var absDir string
	if !filepath.IsAbs(relative) {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to determine current working directory: %w", err)
		}
		absDir = filepath.Join(cwd, relative)
	} else {
		absDir = relative
	}
	return absDir, nil
}

// FilePaths returns the absolute path of the given relative path and a list of files that match the given regex that are in the given directory
func FilePaths(regex, relativeDirPath string) (string, []string, error) {

	re, err := regexp.Compile(regex)
	if err != nil {
		return "", nil, fmt.Errorf("failed to compile regex: %w", err)
	}

	absDir, err := Absolute(relativeDirPath)
	if err != nil {
		return "", nil, err
	}

	result := []string{}
	prefix := absDir + string(filepath.Separator)
	err = filepath.WalkDir(absDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if re.MatchString(path) {
			result = append(result, strings.TrimPrefix(path, prefix))
		}

		return nil
	})
	if err != nil {
		return "", nil, err
	}
	return absDir, result, nil
}
