package traverse

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

// Parameters traverses #/components/parameters
// and individual inline defined parameters
func Parameters(doc *openapi3.T, visitor ParameterRefVisitor) error {

	var err error
	if doc.Components != nil {
		for k, v := range doc.Components.Parameters {
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
		err = ParameterRefs(k, v.Parameters, visitor)
		if err != nil {
			return err
		}

		err = OperationParameterRefs(http.MethodGet, k, v.Get, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodHead, k, v.Head, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodOptions, k, v.Options, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPatch, k, v.Patch, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPost, k, v.Post, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPut, k, v.Put, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodTrace, k, v.Trace, visitor)
		if err != nil {
			return err
		}
	}

	return nil
}

func OperationParameterRefs(method, path string, operation *openapi3.Operation, visitor ParameterRefVisitor) error {
	if operation == nil {
		return nil
	}

	var err error
	for _, v := range operation.Parameters {
		if v.Ref != "" {
			// skip references because they are handled else where
			continue
		}
		if v.Value == nil {
			continue
		}
		name := NameFromOperationParameterRef(method, path, operation.OperationID, v, ParameterNameSuffix)
		err = visitor(name, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParameterRefs(operationName string, list []*openapi3.ParameterRef, visior ParameterRefVisitor) error {
	var err error

	// assumption: all v.Value.Name value in the list are distinct
	for _, v := range list {
		if v.Value == nil {
			continue
		}

		// assumption: final definitions without references do not have a .Ref value set here.
		if v.Ref != "" {
			continue
		}

		name := NameFromParameter(operationName, v.Value, ParameterNameSuffix)
		err = visior(name, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type ParameterRefVisitor func(name string, parameter *openapi3.ParameterRef) error
