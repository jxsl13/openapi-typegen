package testutils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func FilePath(relative string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to get file path")
	}
	if filepath.IsAbs(relative) {
		panic(fmt.Sprintf("%s is an absolute file path, must be relative to the current go source file", relative))
	}
	return filepath.Join(filepath.Dir(file), relative)
}
