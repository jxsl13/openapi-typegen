package testutils

import (
	"fmt"
	"path/filepath"
	"runtime"

	fp "github.com/jxsl13/openapi-typegen/filepath"
)

func FilePath(relative string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to get file path")
	}
	if filepath.IsAbs(relative) {
		panic(fmt.Sprintf("%s is an absolute file path, must be relative to the current go source file", relative))
	}
	abs := filepath.Join(filepath.Dir(file), relative)
	return abs
}

func FilePaths(regex, relativeDirPath string) (string, []string) {
	dir, files, err := fp.FilePaths(regex, FilePath(relativeDirPath))
	if err != nil {
		panic(err)
	}
	return dir, files
}
