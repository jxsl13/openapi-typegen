package strutils

import (
	"regexp"
	"strings"
)

var (
	pathPlaceholder *regexp.Regexp = regexp.MustCompile(`\{([^\{\}]+)\}`)
)

// UnwrapPathParameters unwraps and modifies unwrapped values: {version} -> modify(version)
func UnwrapPathParameters(modifyNew func(string) string, path string) string {
	matches := pathPlaceholder.FindAllStringSubmatch(path, -1)
	if len(matches) == 0 {
		return path
	}

	oldNew := flatten(matches)
	for idx, v := range oldNew {
		if idx%2 == 1 {
			oldNew[idx] = modifyNew(v)
		}
	}

	replacer := strings.NewReplacer(oldNew...)
	return replacer.Replace(path)
}

// UnwrapAllPathParameters unwraps all path placeholders
func UnwrapAllPathParameters(modify func(string) string, path ...string) []string {
	result := make([]string, 0, len(path))
	for _, n := range path {
		result = append(result, UnwrapPathParameters(modify, n))
	}
	return result
}

func flatten(ss [][]string) []string {
	cap := 0
	for _, s := range ss {
		cap += len(s)
	}

	result := make([]string, 0, cap)

	for _, s := range ss {
		result = append(result, s...)
	}
	return result
}
