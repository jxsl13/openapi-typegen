package spec

import "github.com/getkin/kin-openapi/openapi3"

type loadOptions struct {
	validationOptions []openapi3.ValidationOption
}

type Option func(options *loadOptions)

func EnableSchemaFormatValidation() Option {
	return func(options *loadOptions) {
		options.validationOptions = append(options.validationOptions, openapi3.EnableSchemaFormatValidation())
	}
}

func DisableSchemaFormatValidation() Option {
	return func(options *loadOptions) {
		options.validationOptions = append(options.validationOptions, openapi3.DisableSchemaFormatValidation())
	}
}
