package strutils

// Merge concatenates two strings removing duplicate overlaps between joins.
func Merge(a, b string) string {
	o := Overlap(a, b)
	if o > 0 {
		return a + b[o:]
	}
	return a + b
}

// MergeAll concatenates all strings removing duplicate overlaps between joins
// and overlaps across all previous joined strings with the next
// abC + CdeF + Fgh = abCdeFgh
func MergeAll(strs ...string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	var (
		result = strs[0]
		curr   string
	)
	for i := 1; i < len(strs); i++ {
		curr = strs[i]
		result = Merge(result, curr)
	}
	return result
}
