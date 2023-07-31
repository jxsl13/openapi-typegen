package astutils

import (
	"fmt"
	"runtime/debug"
)

var (
	// format for go toolchain pattern matching
	magicCommentFormat = "// Code generated %s DO NOT EDIT"
	// prefixed by the 'by' string
	magicComment  = fmt.Sprintf(magicCommentFormat, "by %s")
	generatorName = ""
)

func SetGeneratorName(name string) {
	generatorName = name
}

func GeneratorComment() string {
	return fmt.Sprintf(magicComment, generatorName)
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		generatorName = "by unknown generator"
	} else if bi.Main.Path == "" {
		generatorName = fmt.Sprintf("test (Go %s)", bi.GoVersion)
	} else {
		generatorName = fmt.Sprintf("%s (%s %s)", bi.Path, bi.Main.Version, bi.Main.Sum)
	}
}
