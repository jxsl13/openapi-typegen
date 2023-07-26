package traverse

import (
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func PathParameters(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if len(doc.Paths) == 0 {
		return nil
	}

	var err error
	for name, path := range doc.Paths {
		err = PathItemParameters(name, path, visitor)
		if err != nil {
			return err
		}
	}
	return nil
}

func PathItemParameters(pathName string, item *openapi3.PathItem, visitor SchemaRefVisitor) error {
	if item == nil {
		return nil
	}

	if item.Ref != "" {
		return nil
	}

	var err error
	for _, param := range item.Parameters {
		if param.Ref != "" || param.Value == nil {
			continue
		}
		paramName := names.Join(
			names.ToTitleTypeName(pathName),
			names.ToTitleTypeName(param.Value.In),
			names.ToTitleTypeName(param.Value.Name),
		)

		for mimeType, mediaType := range param.Value.Content {
			contentSchema := mediaType.Schema
			if contentSchema == nil || contentSchema.Ref != "" || contentSchema.Value == nil {
				continue
			}

			var mtName string
			m := mimetype.Lookup(mimeType)
			if m != nil {
				mtName = m.String()
			} else {
				mtName = mimeType
			}
			mtName = names.ToTitleTypeName(mtName)

			contentTypeHeaderName := names.Join(paramName, mtName, ParameterNameSuffix)
			err = visitor(contentTypeHeaderName, contentSchema)
			if err != nil {
				return err
			}
		}
	}

	err = OperationParameters(pathName, http.MethodConnect, item.Connect, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodDelete, item.Delete, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodGet, item.Get, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodHead, item.Head, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodOptions, item.Options, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodPatch, item.Patch, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodPost, item.Post, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodPut, item.Put, visitor)
	if err != nil {
		return err
	}
	err = OperationParameters(pathName, http.MethodTrace, item.Trace, visitor)
	if err != nil {
		return err
	}

	return nil
}

func OperationParameters(pathName, method string, op *openapi3.Operation, visitor SchemaRefVisitor) error {
	if op == nil {
		return nil
	}

	var (
		err           error
		operationName string
	)

	if op.OperationID != "" {
		operationName = names.ToTitleTypeName(op.OperationID)
	} else {
		operationName = names.Join(
			names.ToTitleTypeName(method),
			names.ToTitleTypeName(pathName),
		)
	}

	for _, param := range op.Parameters {
		if param.Ref != "" || param.Value == nil {
			continue
		}
		paramName := names.Join(
			operationName,
			names.ToTitleTypeName(param.Value.In),
			names.ToTitleTypeName(param.Value.Name),
		)

		for mimeType, mediaType := range param.Value.Content {
			contentSchema := mediaType.Schema
			if contentSchema == nil || contentSchema.Ref != "" || contentSchema.Value == nil {
				continue
			}

			var mtName string
			m := mimetype.Lookup(mimeType)
			if m != nil {
				mtName = m.String()
			} else {
				mtName = mimeType
			}
			mtName = names.ToTitleTypeName(mtName)

			contentTypeHeaderName := names.Join(paramName, mtName, ParameterNameSuffix)
			err = visitor(contentTypeHeaderName, contentSchema)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
