package main

import (
	"embed"
	"net/http"

	"io.github.composeweb/frontend/pages/dashboard"
	"io.github.composeweb/frontend/pages/login"
	"io.github.composeweb/frontend/pages/signup"
	"io.github.composeweb/frontend/pages/view"
	"io.github.composeweb/frontend/template"
)

//go:embed all:components all:pages
var resources embed.FS

func main() {
	template.NewTemplate(resources)

	http.HandleFunc("/signup", signup.SignupPageHandler)
	http.HandleFunc("POST /signup", signup.SignupPostHandler)

	http.HandleFunc("/login", login.LoginPageHandler)
	http.HandleFunc("POST /login", login.LoginPostHandler)

	http.HandleFunc("/dashboard", dashboard.DashboardPageHandler)

	http.HandleFunc("/view/{id}", view.ViewPageHandler)

	http.ListenAndServe(":8080", nil)
}
