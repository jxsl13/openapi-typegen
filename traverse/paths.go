package traverse

import "github.com/getkin/kin-openapi/openapi3"

func Paths(doc openapi3.T, visitor PathVisitor) error {
	if doc.Paths == nil {
		return nil
	}

	var err error
	for k, v := range doc.Paths {
		err = visitor(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type PathVisitor func(pathName string, item *openapi3.PathItem) error
