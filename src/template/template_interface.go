package template

import "github.com/aryadiahmad4689/create-module/src/data"

type TemplateInterface interface {
	Generate(data data.Data) string
}
