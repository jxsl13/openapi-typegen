package traverse

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

// Schemas
func Schemas(doc *openapi3.T, visitor SchemaRefVisitor) error {
	var err error
	if doc.Components != nil {
		for name, schema := range doc.Components.Schemas {

			if schema.Ref != "" || schema.Value == nil {
				continue
			}
			// only traverse on-references
			err = visitor(names.ToTitleTypeName(name), schema)
			if err != nil {
				return err
			}
		}

		for name, param := range doc.Components.Parameters {
			if param.Ref != "" || param.Value == nil {
				continue
			}

			if param.Value.Schema != nil {
				err = visitor(names.ToTitleTypeName(name), param.Value.Schema)
				if err != nil {
					return err
				}
			}

			for mimeType, mediaType := range param.Value.Content {
				if mediaType.Schema.Ref != "" || mediaType.Schema.Value == nil {
					continue
				}

				m := mimetype.Lookup(mimeType)
				if m == nil {
					err = visitor(NameFromComponentRequestBodyMediaType(
						name,
						mimeType,
						RequestBodyNameSuffix,
					), mediaType.Schema)
					if err != nil {
						return err
					}
				} else {
					err = visitor(NameFromComponentRequestBodyMediaType(
						name,
						m.String(),
						RequestBodyNameSuffix,
					), mediaType.Schema)
					if err != nil {
						return err
					}
				}
			}

		}

		for name, header := range doc.Components.Headers {
			if header.Ref != "" || header.Value == nil {
				continue
			}

			if header.Value.Schema == nil || header.Value.Schema.Ref != "" || header.Value.Schema.Value == nil {
				continue
			}

			err = visitor(names.ToTitleTypeName(name), header.Value.Schema)
			if err != nil {
				return err
			}
		}

		for name, req := range doc.Components.RequestBodies {
			if req.Ref != "" || req.Value == nil {
				continue
			}

			for mimeType, mediaType := range req.Value.Content {
				if mediaType.Schema.Ref != "" || mediaType.Schema.Value == nil {
					continue
				}

				m := mimetype.Lookup(mimeType)
				if m == nil {
					err = visitor(NameFromComponentRequestBodyMediaType(
						name,
						mimeType,
						RequestBodyNameSuffix,
					), mediaType.Schema)
					if err != nil {
						return err
					}
				} else {
					err = visitor(NameFromComponentRequestBodyMediaType(
						name,
						m.String(),
						RequestBodyNameSuffix,
					), mediaType.Schema)
					if err != nil {
						return err
					}
				}
			}
		}

		for name, resp := range doc.Components.Responses {
			if resp.Ref != "" || resp.Value == nil {
				continue
			}

			respName := names.ToTitleTypeName(name)

			for mimeType, mediaType := range resp.Value.Content {
				if mediaType.Schema.Ref != "" || mediaType.Schema.Value == nil {
					continue
				}

				var respBodyName string
				if len(resp.Value.Content) == 1 {
					respBodyName = respName
				} else {

					m := mimetype.Lookup(mimeType)
					if m == nil {
						respBodyName = NameFromComponentRequestBodyMediaType(
							respName,
							mimeType,
							RequestBodyNameSuffix,
						)
					} else {
						respBodyName = NameFromComponentRequestBodyMediaType(
							respName,
							m.String(),
							RequestBodyNameSuffix,
						)
					}
				}

				err = visitor(respBodyName, mediaType.Schema)
				if err != nil {
					return err
				}
			}

			for name, header := range resp.Value.Headers {
				if header.Ref != "" || header.Value == nil {
					continue
				}

				if header.Value.Schema.Ref != "" || header.Value.Schema.Value == nil {
					continue
				}

				headerName := names.Join(
					respName,
					names.ToTitleTypeName(name),
					names.ToTitleTypeName(header.Value.Name),
					names.ToTitleTypeName(HeaderNameSuffix),
				)

				// header is only a parameter wrapped in a header type
				// we can iterate over it like we do with a parameter
				err = visitor(headerName, header.Value.Schema)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

type SchemaRefVisitor func(name string, schema *openapi3.SchemaRef) error
