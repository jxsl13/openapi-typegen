package traverse

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var (
	RequestType   = "request"
	ResponseType  = "response"
	ParameterType = "parameter"
	SchemaType    = "schema"

	TypeKey      = "type"
	NameKey      = "name"
	InKey        = "in"
	PathKey      = "path"
	MethodKey    = "method"
	OperationKey = "operation"
	ContentKey   = "content"
	StatusKey    = "status"
)

type SchemaVisitor func(schemaRef *openapi3.SchemaRef, levelNames map[string][]string) error

// Document traverses the given document and calls the visitor for each schema.
func Document(t *openapi3.T, visitor SchemaVisitor, levelNames ...map[string][]string) error {
	if t == nil {
		return nil
	}

	var err error
	for pathName, pathItem := range t.Paths {
		if pathItem == nil {
			continue
		}
		if pathItem.Ref != "" {
			continue
		}

		err = PathItem(pathItem, visitor, add(merge(levelNames...), PathKey, pathName))
		if err != nil {
			return err
		}
	}

	// traverse components
	if t.Components != nil {
		err = Components(t.Components, visitor, merge(levelNames...))
		if err != nil {
			return err
		}
	}

	return nil
}

// merge any number of maps into a new map
func merge(maps ...map[string][]string) map[string][]string {
	size := 0
	for _, m := range maps {
		size += len(m)
	}

	m := make(map[string][]string, size)
	for _, m1 := range maps {
		for k, vs := range m1 {
			m[k] = append(m[k], vs...)
		}
	}
	return m
}

// add clones the map and adds all new
// empty key values are not added
func add(m map[string][]string, keyValue ...string) map[string][]string {
	if len(keyValue)%2 != 0 {
		panic("keyValue must be even")
	}

	m2 := make(map[string][]string, len(m)+len(keyValue)/2)
	for k, v := range m {
		m2[k] = append(m2[k], v...)
	}

	for i := 0; i < len(keyValue); i += 2 {
		if keyValue[i+1] != "" {
			m2[keyValue[i]] = append(m2[keyValue[i]], keyValue[i+1]) // only add non empty keys
		}
	}

	return m2
}
