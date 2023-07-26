package traverse

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
)

func ComponentRequestBodies(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil {
		return nil
	}

	var err error
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
				err = visitor(NameFromComponentMediaType(
					name,
					mimeType,
					RequestBodyNameSuffix,
				), mediaType.Schema)
				if err != nil {
					return err
				}
			} else {
				err = visitor(NameFromComponentMediaType(
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
	return nil
}
