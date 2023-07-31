package strutils

import "strings"

func ToTitle(name string) string {
	//lint:ignore SA1019 for our use case this function is enough, as we only work with alphanumeric characters
	return strings.Title(name)

}
