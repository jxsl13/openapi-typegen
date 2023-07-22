package traverse

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func Parameters(doc *openapi3.T, visitor ParameterRefVisitor) error {

	parameters := doc.Components.Parameters
	var err error
	for k, v := range parameters {
		if v.Value == nil {
			panic(fmt.Sprintf("component parameter %q is nil", k))
		}
		err = visitor(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

type ParameterRefVisitor func(name string, parameter *openapi3.ParameterRef) error
