package traverse

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func ComponentResponses(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil {
		return nil
	}

	var err error
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
					respBodyName = NameFromComponentMediaType(
						respName,
						mimeType,
						RequestBodyNameSuffix,
					)
				} else {
					respBodyName = NameFromComponentMediaType(
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

	return nil
}
