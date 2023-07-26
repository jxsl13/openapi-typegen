package traverse

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jxsl13/openapi-typegen/names"
)

func ComponentParameters(doc *openapi3.T, visitor SchemaRefVisitor) error {
	if doc.Components == nil {
		return nil
	}
	var err error
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
