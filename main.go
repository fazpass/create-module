package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fazpass/create-module/src/data"
	mains "github.com/fazpass/create-module/src/main"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var moduleName = flag.String("module_name", "", "Ini adalah flag module name exmaples: -module_name=test")
	var rootDir = flag.String("root_dir", "", "Ini adalah flag root directory exmaples: -root_dir=src/halo")
	var moduleInit = flag.String("module_init", "", "Ini adalah flag module init exmaples: -module_init=datz/common")
	flag.Parse()

	if *moduleName == "" {
		fmt.Println("module name is empty")
		os.Exit(1)
	}
	var modInit string
	var rootDirectory string
	if *moduleInit == "" {
		modInit = os.Getenv("MODULE_INIT")
	} else {
		modInit = *moduleInit
	}

	if *rootDir == "" {
		rootDirectory = os.Getenv("ROOT_DIR")
	} else {
		rootDirectory = *rootDir
	}

	var err = mains.InitMain().Create(data.MainData{
		ModuleName:       *moduleName,
		RootDir:          rootDirectory,
		ProjectDirectory: data.ProjectDirectory,
		ModInit:          modInit,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
