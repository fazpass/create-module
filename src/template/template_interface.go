package template

import "github.com/fazpass/create-module/src/data"

type TemplateInterface interface {
	Generate(data data.Data) string
}
