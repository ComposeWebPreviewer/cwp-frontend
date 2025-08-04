package view

import (
	"encoding/json"
	"io"
	"net/http"

	"io.github.composeweb/frontend/template"
)

type MyResponse struct {
	Status  string `json:"status"`
Code string `json:"code"`
Wasm string `json:"wasm"`
Js string `json:"js"`
}

func ViewPageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	res, _ := http.Get("https://20s5mqesgj.execute-api.ap-south-1.amazonaws.com/Prod/?id=" + id)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response MyResponse
	json.Unmarshal(body, &response)
	template.Render(w, "view.html", struct {
		Js string
		Wasm string
	} {
		Js:   response.Js,
			Wasm: response.Wasm,
	})
}
