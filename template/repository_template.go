package template

import "github.com/aryadiahmad4689/create-module/utils"

func GenerateRepositoryTemplate(data utils.Data) string {
	var template = `package repository

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
	return utils.Parse("repository.go", template, data)

}
