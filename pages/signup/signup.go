package signup

import (
	"net/http"

	"io.github.composeweb/frontend/template"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	template.Render(w, "signup.html", nil)
}
