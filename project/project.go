package project

import (
	"errors"
	"fmt"
	"os"

	"github.com/aryadiahmad4689/create-module/template"
	"github.com/aryadiahmad4689/create-module/utils"
	_ "github.com/joho/godotenv/autoload"
)

var projectDirectory = []string{"repository", "usecase", "endpoint", "transport/http/handler"}

func Create(moduleName string) error {
	if moduleName == "" {
		return errors.New("module name is empty")
	}
	err := utils.CreateFolder(os.Getenv("ROOT_DIR") + moduleName)
	if err != nil {
		return err
	}
	err = utils.CreateSubFolder(projectDirectory, os.Getenv("ROOT_DIR")+moduleName)
	if err != nil {
		return err
	}
	data := utils.Data{
		ModuleName: moduleName,
		ModInit:    os.Getenv("MODULE_INIT"),
	}
	fmt.Println(data, os.Getenv("MODULE_INIT"), os.Getenv("ROOT_DIR"))
	// create module
	utils.CreateFile(os.Getenv("ROOT_DIR")+moduleName+"/module.go", []byte(template.GenerateModuleTemplate(data)))
	// create Repository
	utils.CreateFile(os.Getenv("ROOT_DIR")+moduleName+"/repository/repository.go", []byte(template.GenerateRepositoryTemplate(data)))
	// create UseCase
	utils.CreateFile(os.Getenv("ROOT_DIR")+moduleName+"/usecase/usecase.go", []byte(template.GenerateUsecaseTemplate(data)))
	// create Endpoint
	utils.CreateFile(os.Getenv("ROOT_DIR")+moduleName+"/endpoint/endpoint.go", []byte(template.GenerateEndpointTemplate(data)))
	// create Transport
	utils.CreateFile(os.Getenv("ROOT_DIR")+moduleName+"/transport/http/handler/handler.go", []byte(template.GenerateTransportTemplate(data)))
	return nil
}
