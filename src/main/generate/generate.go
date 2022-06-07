package generate

import (
	"github.com/fazpass/create-module/src/data"
	"github.com/fazpass/create-module/src/template"
	"github.com/fazpass/create-module/src/utils"
)

type GenerateInterface interface {
	Excute(data data.MainData) error
}

type Generate struct {
	template template.TemplateInterface
}

func Init() GenerateInterface {
	return &Generate{
		template: template.Init(),
	}
}

func (gen *Generate) Excute(da data.MainData) error {

	var dataGenerate = []data.DataGenerate{
		{
			ModuleName:    da.ModuleName,
			FileName:      "module.go",
			Template:      data.TemplateModule,
			DirectoryFile: "/module.go",
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},

		{
			ModuleName:    da.ModuleName,
			FileName:      "repository.go",
			DirectoryFile: "/repository/repository.go",
			Template:      data.TemplateRepository,
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},
		{
			ModuleName:    da.ModuleName,
			FileName:      "usecase.go",
			DirectoryFile: "/usecase/usecase.go",
			Template:      data.TemplateUsecase,
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},

		{
			ModuleName:    da.ModuleName,
			FileName:      "endpoint.go",
			DirectoryFile: "/endpoint/endpoint.go",
			Template:      data.TemplateEndPoint,
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},
		{
			ModuleName:    da.ModuleName,
			FileName:      "handler.go",
			DirectoryFile: "/transport/http/handler/handler.go",
			Template:      data.TemplateTransport,
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},
		{
			ModuleName:    da.ModuleName,
			FileName:      "router.go",
			DirectoryFile: "/transport/http/router.go",
			Template:      data.TemplateTransport,
			ModuleInit:    da.ModInit,
			RootDir:       da.RootDir,
		},
	}

	for _, v := range dataGenerate {
		var data = data.Data{
			ModuleName: v.ModuleName,
			ModInit:    v.ModuleInit,
			Template:   v.Template,
			FileName:   v.FileName,
		}

		var err = utils.CreateFile(v.RootDir+v.ModuleName+v.DirectoryFile, []byte(gen.template.Generate(data)))
		if err != nil {
			return err
		}
	}
	return nil
}
