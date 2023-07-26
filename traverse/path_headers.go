package traverse

import (
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func PathHeaders(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if len(doc.Paths) == 0 {
		return nil
	}

	var err error
	for name, path := range doc.Paths {
		err = PathItemHeaders(name, path, visitor)
		if err != nil {
			return err
		}
	}
	return nil
}

func PathItemHeaders(pathName string, item *openapi3.PathItem, visitor SchemaRefVisitor) error {
	if item == nil {
		return nil
	}

	if item.Ref != "" {
		return nil
	}
	var err error
	err = OperationHeaders(pathName, http.MethodConnect, item.Connect, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodDelete, item.Delete, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodGet, item.Get, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodHead, item.Head, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodOptions, item.Options, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodPatch, item.Patch, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodPost, item.Post, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodPut, item.Put, visitor)
	if err != nil {
		return err
	}
	err = OperationHeaders(pathName, http.MethodTrace, item.Trace, visitor)
	if err != nil {
		return err
	}

	return nil
}

func OperationHeaders(path, method string, op *openapi3.Operation, visitor SchemaRefVisitor) error {
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
			names.ToTitleTypeName(path),
		)
	}

	responses := op.Responses
	for respName, resp := range responses {
		if resp.Ref != "" || resp.Value == nil {
			continue
		}

		for headerName, header := range resp.Value.Headers {
			if header.Ref != "" || header.Value == nil {
				continue
			}

			schema := header.Value.Schema

			if schema == nil || schema.Ref != "" || schema.Value == nil {
				continue
			}

			responseHeaderName := names.Join(
				operationName,
				names.ToTitleTypeName(respName),
				names.ToTitleTypeName(headerName),
				names.ToTitleTypeName(header.Value.Name),
			)

			err = visitor(names.Join(responseHeaderName, HeaderNameSuffix), schema)
			if err != nil {
				return err
			}

			content := header.Value.Content
			for mimeType, mediaType := range content {
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

				contentTypeHeaderName := names.Join(responseHeaderName, mtName, HeaderNameSuffix)
				err = visitor(contentTypeHeaderName, contentSchema)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}
