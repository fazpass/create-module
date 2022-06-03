package main

import (
	"log"
	"os"

	"github.com/aryadiahmad4689/create-module/project"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please define your name module")
	}
	err := project.Create(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
