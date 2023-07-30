package names

import (
	"regexp"
	"strings"
)

var (
	nonAlphaNum     *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9]`)
	numericPrefix   *regexp.Regexp = regexp.MustCompile(`^[0-9]+`)
	pathPlaceholder *regexp.Regexp = regexp.MustCompile(`\{([^\{\}]+)\}`)
)

// ToTitleTypeName first makes the string title cased and then
// transforms it to a Go compliant type name (identifier)
func ToTitleTypeName(name string) string {
	return ToTypeName(ToTitle(name))
}

// ToTypeName removes all non alpha numeric characters
// Because OperationIDs are supposed to be universally unique,
// we do not add anything in front or at the end (for now).
// Tho, identifiers starting with an integer will get a N prefix.
func ToTypeName(name string) string {
	name = PrefixInteger(name)
	return toTypeName(name)
}

func toTypeName(name string) string {
	// some api specs use weird names
	// we make plural from array types
	if strings.HasSuffix(name, "[]") {
		name = Join(name[:len(name)-2], "s")
	}

	name = UnwrapPathPlaceholder(ToTitle, name)

	return nonAlphaNum.ReplaceAllString(name, "")
}

func ToTitle(name string) string {
	//lint:ignore SA1019 for our use case this function is enough, as we only work with alphanumeric characters
	return strings.Title(name)

}

func PrefixInteger(name string) string {
	if numericPrefix.MatchString(name) {
		name = "N" + name
	}
	return name
}

func JoinTitleTypeNames(names ...string) string {
	titledTypes := make([]string, 0, len(names))
	for _, name := range names {
		titledTypes = append(titledTypes, toTypeName(ToTitle(name)))
	}

	return PrefixInteger(Join(titledTypes...))
}

func Append(name string, names ...string) string {
	return strings.Join(append([]string{name}, names...), "")
}

func Ignore(ignored string, names ...string) []string {
	result := make([]string, 0, len(names))
	for _, name := range names {
		if !strings.EqualFold(ignored, name) {
			result = append(result, name)
		}
	}
	return result
}

// Join concatenates all strings removing duplicate overlaps between joins
// and overlaps across all previous joined strings with the next
// abC + CdF + Fgh = abCdeFgh
func Join(names ...string) string {
	if len(names) == 0 {
		return ""
	} else if len(names) == 1 {
		return names[0]
	}

	var (
		result = names[0]
		curr   string
	)
	for i := 1; i < len(names); i++ {
		curr = names[i]
		result = Merge(result, curr)
	}
	return result
}

// Unwraps and modifies unwrapped values: {version} -> modify(version)
func UnwrapPathPlaceholder(modifyNew func(string) string, name string) string {
	matches := pathPlaceholder.FindAllStringSubmatch(name, -1)
	if len(matches) == 0 {
		return name
	}

	oldNew := flatten(matches)
	for idx, v := range oldNew {
		if idx%2 == 1 {
			oldNew[idx] = modifyNew(v)
		}
	}

	replacer := strings.NewReplacer(oldNew...)
	return replacer.Replace(name)
}

// UnwrapAllPathPlaceholders unwraps all path placeholders
func UnwrapAllPathPlaceholders(modify func(string) string, name ...string) []string {
	result := make([]string, 0, len(name))
	for _, n := range name {
		result = append(result, UnwrapPathPlaceholder(modify, n))
	}
	return result
}

// a and b are ordered and merged with a's suffix and b's prefix
// to deduplicate names
func Merge(a, b string) string {
	o := Overlap(a, b)
	if o > 0 {
		return a + b[o:]
	}
	return a + b
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
