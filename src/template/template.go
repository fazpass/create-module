package template

import (
	"github.com/fazpass/create-module/src/data"
	"github.com/fazpass/create-module/src/utils"
)

type Template struct {
}

func Init() TemplateInterface {
	return &Template{}
}

func (templ *Template) Generate(data data.Data) string {
	return utils.Parse(data.FileName, data.Template, data)
}
