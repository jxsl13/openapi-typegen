package strutils

func Size(ss ...string) int {
	size := 0
	for i := 0; i < len(ss); i++ {
		size += len(ss[i])
	}
	return size
}
