package main

import (
	"embed"
	"net/http"

	"io.github.composeweb/frontend/pages/home"
	"io.github.composeweb/frontend/pages/view"
	"io.github.composeweb/frontend/template"
)

//go:embed all:components all:pages
var resources embed.FS

func main() {
	template.NewTemplates(resources)

	http.HandleFunc("/", home.HomePageHandler)
	http.HandleFunc("/view/{id}", view.ViewPageHandler)

	http.ListenAndServe(":8080", nil)
}
