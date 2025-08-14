package dashboard

import (
	"net/http"

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
	}

	

	template.Render(w, "dashboard.html", nil)
}
