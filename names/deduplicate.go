package names

import "strings"

func Deduplicate(names []string) []string {
	dups := make(map[int]bool, len(names)-1)
	for i := range names {
		for j := range names {
			if i == j {
				continue
			}
			ni := strings.ToLower(names[i])
			nj := strings.ToLower(names[j])
			if strings.Contains(ni, nj) {
				dups[j] = true
			} else if strings.Contains(nj, ni) {
				dups[i] = true
			}
		}
	}

	result := make([]string, 0, len(names)-len(dups))
	for idx, name := range names {
		if !dups[idx] {
			result = append(result, name)
		}
	}

	return result
}
