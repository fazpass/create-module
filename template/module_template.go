package template

import (
	"github.com/aryadiahmad4689/create-module/utils"
)

func GenerateModuleTemplate(data utils.Data) string {
	var template = `package module

	import (
		"{{.ModInit}}/src/app"
		modulebase "{{.ModInit}}/src/app/module"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/endpoint"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/repository"
		transporthttp "{{.ModInit}}/src/modules/{{.ModuleName}}/transport/http"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/usecase"
		"github.com/go-chi/chi"
	)

	type Module struct {
		usecase    usecase.UseCaseInterface
		endpoint   endpoint.EndpointInterface
		repository *repository.Repository
		httpRouter *chi.Mux
	}
	
	func Init(app *app.App) modulebase.ModuleInterface {
		var (
			repository = repository.Init(app)
			usecase    = usecase.Init(app, repository)
			endpoint   = endpoint.Init(app, usecase)
			httpRouter = transporthttp.Init(app, endpoint)
		)
	
		return &Module{
			repository: repository,
			usecase:    usecase,
			endpoint:   endpoint,
			httpRouter: httpRouter,
		}
	}
	
	func (module *Module) GetHttpRouter() *chi.Mux {
		return module.httpRouter
	}
	`
	return utils.Parse("module.go", template, data)

}
