package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
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
