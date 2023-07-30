package options

import (
	"github.com/jxsl13/openapi-typegen/names"
	"github.com/jxsl13/openapi-typegen/traverse"
)

func UniqueName(levels map[string][]string) string {

	var parts []string
	if len(levels[traverse.OperationKey]) != 0 {
		// using operation id
		parts = names.SplitAll(levels[traverse.OperationKey]...)
		parts = append(parts, levels[traverse.NameKey]...) // we don't want to merge names with duplicates
	} else {
		// no operation id, using other keys
		parts = names.SplitAll(levels[traverse.MethodKey]...)

		pathName := names.JoinTitleTypeNames(names.UnwrapAllPathPlaceholders(names.ToTitle, levels[traverse.PathKey]...)...)
		pathName = names.Append(pathName, levels[traverse.NameKey]...)
		parts = append(parts, pathName) // we don't want to merge names with duplicates
	}

	parts = append(parts, names.SplitAll(levels[traverse.InKey]...)...)

	parts = append(parts, names.SplitAll(names.MimeTypes(levels[traverse.ContentKey])...)...)
	parts = append(parts, names.SplitAll(names.StatusCodes(levels[traverse.StatusKey])...)...)

	parts = append(parts, names.SplitAll(names.Ignore(traverse.SchemaType, levels[traverse.TypeKey]...)...)...)

	return names.JoinTitleTypeNames(parts...)

}
