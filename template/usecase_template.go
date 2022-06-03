package template

import "github.com/aryadiahmad4689/create-module/utils"

func GenerateUsecaseTemplate(data utils.Data) string {
	var template = `package usecase

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
	return utils.Parse("usecase.go", template, data)

}
