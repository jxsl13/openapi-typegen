package strutils

import "strings"

// Deduplicate removes duplicate string that are either equal or contain each other.
// The string contained in the other string will be removed.
func Deduplicate(names []string) []string {
	result := make([]string, 0, len(names))
	for i := 0; i < len(names); i++ {
		// we look at ni whether to add it to the result
		ni := strings.ToLower(names[i])
		dup := false
		for j := 0; j < len(names); j++ {
			if i == j {
				continue
			}
			nj := strings.ToLower(names[j])
			// other string contains current string
			if strings.Contains(nj, ni) && ni != nj {
				dup = true
				break
			}
		}
		if !dup {
			dup = false
			for _, r := range result {
				if strings.Contains(strings.ToLower(r), ni) {
					dup = true
					break
				}
			}
			if !dup {
				result = append(result, names[i])
			}

		}
	}

	return result

}
