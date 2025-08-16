package dashboard

import (
	"fmt"
	"net/http"

	"io.github.composeweb/frontend/api"
	"io.github.composeweb/frontend/components/composablecard"
	"io.github.composeweb/frontend/template"
)

type dashboardPageData struct {
	Cards []composablecard.ComposableCard
}

func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	authorization, err := r.Cookie("Authorization")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	codespaces, err := api.GetCodespaces(authorization.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var cards []composablecard.ComposableCard
	for _, codespace := range codespaces.Codespaces {
		cards = append(cards, composablecard.ComposableCard{
			Title:     "Codespace",
			ImageURL:  "https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp",
			ImageAlt:  "Codespace Image",
			ButtonURL: fmt.Sprintf("/view/?id=%s&uid=%s", codespace.Id, codespace.Uid),
		})
	}

	template.Render(w, "dashboard.html", dashboardPageData{
		Cards: cards,
	})
}
