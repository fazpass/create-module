package data

type Data struct {
	ModuleName string
	ModInit    string
	Template   string
	FileName   string
}

type MainData struct {
	ModuleName       string
	ProjectDirectory []string
	RootDir          string
	ModInit          string
}

// ini di gunakan untuk menggenerate template
type DataGenerate struct {
	ModuleName    string
	DirectoryFile string
	RootDir       string
	ModuleInit    string
	Template      string
	FileName      string
}

var (
	ProjectDirectory = []string{"repository", "usecase", "endpoint", "transport/http/handler"}

	TemplateUsecase = `package usecase

	import (
	"{{.ModInit}}/src/app"
	"{{.ModInit}}/src/modules/{{.ModuleName}}/repository"
	)

type (
	UseCaseInterface interface {
	
	}

	UseCase struct {
		app  *app.App
		repo *repository.Repository
	}
)

func Init(app *app.App, repo *repository.Repository) UseCaseInterface {
	return &UseCase{
		app:  app,
		repo: repo,
	}
}
`
	TemplateEndPoint = `package endpoint

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
	TemplateRepository = `package repository

	import (
		"{{.ModInit}}/src/app"
	)
	
	type (
		Repository struct {
		}
	)
	
	func Init(app *app.App) *Repository {
		return &Repository{
		}
	}
	`
	TemplateModule = `package module

	import (
		"{{.ModInit}}/src/app"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/endpoint"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/repository"
		transporthttp "{{.ModInit}}/src/modules/{{.ModuleName}}/transport/http"
		"{{.ModInit}}/src/modules/{{.ModuleName}}/usecase"
		modulebase "github.com/fazpass/goliath/module"
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
	TemplateTransport = `package handler

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
	TemplateRoute = `package http

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
)
