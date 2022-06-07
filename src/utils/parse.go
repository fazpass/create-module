package utils

import (
	"bytes"
	"html/template"
	"log"

	"github.com/fazpass/create-module/src/data"
)

func Parse(name, tplt string, t data.Data) string {
	tmpl, err := template.New(name).Parse(tplt)
	if err != nil {
		log.Fatal(err.Error())
	}
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, t); err != nil {
		log.Fatal(err.Error())
	}
	return buffer.String()
}
