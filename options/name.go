package options

import (
	"github.com/jxsl13/openapi-typegen/strutils"
	"github.com/jxsl13/openapi-typegen/traverse"
)

func UniqueName(levels map[string][]string) string {

	var parts []string
	if len(levels[traverse.OperationKey]) != 0 {
		// using operation id
		parts = strutils.SplitAll(levels[traverse.OperationKey]...)
		parts = append(parts, levels[traverse.NameKey]...) // we don't want to merge strutils with duplicates
	} else {
		// no operation id, using other keys
		parts = strutils.SplitAll(levels[traverse.MethodKey]...)

		pathName := strutils.MergeAllTypeNames(strutils.UnwrapAllPathParameters(strutils.ToTitle, levels[traverse.PathKey]...)...)
		pathName = strutils.Append(pathName, levels[traverse.NameKey]...)
		parts = append(parts, pathName) // we don't want to merge strutils with duplicates
	}

	parts = append(parts, strutils.SplitAll(levels[traverse.InKey]...)...)

	parts = append(parts, strutils.SplitAll(strutils.MimeTypeNames(levels[traverse.ContentKey])...)...)
	parts = append(parts, strutils.SplitAll(strutils.StatusCodes(levels[traverse.StatusKey])...)...)

	parts = append(parts, strutils.SplitAll(strutils.Ignore(traverse.SchemaType, levels[traverse.TypeKey]...)...)...)

	return strutils.MergeAllTypeNames(parts...)

}
