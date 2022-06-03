package template

import "github.com/aryadiahmad4689/create-module/utils"

func GenerateRouteTemplate(data utils.Data) string {
	var template = `package http

	import (
		"fmt"
		"{{.ModInit}}/src/app"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/endpoint"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/transport/http/handler"
		"github.com/go-chi/chi"
	
	)
	
	func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
		var (
			router      = chi.NewRouter()
			h           = handler.InitHandler(app, endpoint)
		)
		
		fmt.Print(h)
	
		return router
	}
	`
	return utils.Parse("router.go", template, data)

}
