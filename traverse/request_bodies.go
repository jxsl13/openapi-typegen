package traverse

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

var RequestBodyNameSuffix = "Request"

// Parameters traverses #/components/parameters
// and individual inline defined parameters
func RequestBodies(doc *openapi3.T, visitor RequestBodyRefVisitor) error {

	var err error

	if doc.Components != nil {
		for k, v := range doc.Components.RequestBodies {
			if v.Ref != "" {
				continue
			}
			if v.Value == nil {
				continue
			}
			err = visitor(k, v)
			if err != nil {
				return err
			}
		}
	}

	for k, v := range doc.Paths {
		if v.Ref != "" {
			// assumption: seemingly it is possible to reference complete path implementations
			// we only want to iterate over local definitions. Global schemas should be handled elsewhere.
			// TODO: check if path references can be defined globally
			continue
		}

		err = OperationRequestBodyRefs(http.MethodGet, k, v.Get, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodHead, k, v.Head, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodOptions, k, v.Options, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodPatch, k, v.Patch, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodPost, k, v.Post, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodPut, k, v.Put, visitor)
		if err != nil {
			return err
		}
		err = OperationRequestBodyRefs(http.MethodTrace, k, v.Trace, visitor)
		if err != nil {
			return err
		}

	}

	return nil
}

func OperationRequestBodyRefs(method, path string, operation *openapi3.Operation, visitor RequestBodyRefVisitor) error {
	if operation == nil {
		return nil
	}

	var err error
	if operation.RequestBody == nil {
		return nil
	} else if operation.RequestBody.Ref != "" {
		// skip references
		return nil
	}

	name := NameFromOperation(method, path, operation, RequestBodyNameSuffix)
	err = visitor(name, operation.RequestBody)
	if err != nil {
		return err
	}

	return nil
}

type RequestBodyRefVisitor func(name string, request *openapi3.RequestBodyRef) error
