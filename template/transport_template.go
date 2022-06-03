package template

import "github.com/aryadiahmad4689/create-module/utils"

func GenerateTransportTemplate(data utils.Data) string {
	var template = `package handler

	import (
		"{{.ModInit}}/src/app"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/endpoint"
	)
	
	type Handler struct {
		app      *app.App
		endpoint endpoint.EndpointInterface
	}
	
	func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
		var handler = &Handler{
			app:      app,
			endpoint: endpoint,
		}
		return handler
	}	
	`
	return utils.Parse("handler.go", template, data)

}
