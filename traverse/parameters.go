package traverse

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

var ParameterNameSuffix = "Parameter"

// Parameters traverses #/components/parameters
// and individual inline defined parameters
func Parameters(doc *openapi3.T, visitor ParameterRefVisitor) error {
	if doc.Components == nil {
		return nil
	}

	var (
		parameters = doc.Components.Parameters
		err        error
	)

	for k, v := range parameters {
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

	for k, v := range doc.Paths {
		if v.Ref != "" {
			// assumption: seemingly it is possible to reference complete path implementations
			// we only want to iterate over local definitions. Global schemas should be handled elsewhere.
			// TODO: check if path references can be defined globally
			continue
		}
		err = ParameterRefs(names.Join(names.ToTitleTypeName(k), ParameterNameSuffix), v.Parameters, visitor)
		if err != nil {
			return err
		}

		err = OperationParameterRefs(http.MethodGet, v.Get, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodHead, v.Head, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodOptions, v.Options, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPatch, v.Patch, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPost, v.Post, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodPut, v.Put, visitor)
		if err != nil {
			return err
		}
		err = OperationParameterRefs(http.MethodTrace, v.Trace, visitor)
		if err != nil {
			return err
		}

	}

	return nil
}

func OperationParameterRefs(method string, operation *openapi3.Operation, visitor ParameterRefVisitor) error {
	if operation == nil {
		return nil
	}

	var (
		name string
		err  error
	)
	for _, v := range operation.Parameters {
		if v.Ref != "" {
			// skip references because they are handled else where
			continue
		}
		if v.Value == nil {
			continue
		}

		// TODO: make this name construction modifiable with a custom name construction function
		if operation.OperationID != "" {
			name = names.Join(names.ToTitleTypeName(operation.OperationID), names.ToTitleTypeName(v.Value.Name))
		} else {
			name = names.Join(names.ToTitle(method), names.ToTitleTypeName(v.Value.Name))
		}

		name = names.Join(name, ParameterNameSuffix)

		err = visitor(name, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParameterRefs(namePrefix string, list []*openapi3.ParameterRef, visior ParameterRefVisitor) error {
	var (
		name string
		err  error
	)

	// assumption: all v.Value.Name value in the list are distinct
	for _, v := range list {
		if v.Value == nil {
			continue
		}

		// assumption: final definitions without references do not have a .Ref value set here.
		if v.Ref != "" {
			continue
		}

		// TODO: make this name construction modifiable with a custom name construction function
		name = names.Join(names.ToTitleTypeName(namePrefix), names.ToTitleTypeName(v.Value.Name))
		name = names.Join(name, ParameterNameSuffix)

		err = visior(name, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type ParameterRefVisitor func(name string, parameter *openapi3.ParameterRef) error
