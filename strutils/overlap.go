package strutils

import "strings"

// Overlap returns the maximum overlap of two strings.
// The end of the first string is compared to the beginning of the second string.
// Order of the strings is important.
func Overlap(a, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	maxOverlap := len(b)
	if len(a) < len(b) {
		maxOverlap = len(a)
	}

	a = a[len(a)-maxOverlap:]
	b = b[:maxOverlap]

	var (
		currentMaxOverlap = 0

		aa     string = a[:]
		suffix string
	)

	for i := strings.IndexByte(aa, b[0]); i >= 0; i = strings.IndexByte(aa, b[0]) {
		suffix = aa[i:]
		if strings.HasPrefix(b, suffix) {
			currentMaxOverlap = len(suffix)
			break
		}

		aa = aa[1:]
	}

	return currentMaxOverlap
}
