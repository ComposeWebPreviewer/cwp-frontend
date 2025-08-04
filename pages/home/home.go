package home

import (
	"net/http"

	"io.github.composeweb/frontend/template"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "home.html", nil)
}
