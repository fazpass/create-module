package template

import (
	"github.com/aryadiahmad4689/create-module/utils"
)

func GenerateEndpointTemplate(data utils.Data) string {
	var template = `package endpoint

import (

	"{{.ModInit}}/src/app"
	"{{.ModInit}}/src/modules/{{.ModuleName}}/usecase"
)

type (
	EndpointInterface interface {
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
	`

	return utils.Parse("endpoint.go", template, data)
}
