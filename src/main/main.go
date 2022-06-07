package mains

import (
	"errors"
	"fmt"

	"github.com/fazpass/create-module/src/data"
	"github.com/fazpass/create-module/src/main/generate"
	"github.com/fazpass/create-module/src/utils"

	_ "github.com/joho/godotenv/autoload"
)

type MainInterface interface {
	Create(data.MainData) error
}

type Main struct {
	generate generate.GenerateInterface
}

func InitMain() MainInterface {
	return &Main{
		generate: generate.Init(),
	}
}

func (m *Main) Create(data data.MainData) error {
	if data.ModuleName == "" {
		return errors.New("module name is empty")
	}
	err := utils.CreateFolder(data.RootDir + data.ModuleName)
	if err != nil {
		return err
	}
	err = utils.CreateSubFolder(data.ProjectDirectory, data.RootDir+data.ModuleName)
	if err != nil {
		return err
	}
	// excecute template
	err = m.generate.Excute(data)
	if err != nil {
		return err
	}

	fmt.Println("Create Success Module: " + data.ModuleName)
	return nil
}
