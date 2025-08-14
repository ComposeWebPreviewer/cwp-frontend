package login

import (
	"fmt"
	"net/http"
	"time"

	"io.github.composeweb/frontend/api"
	"io.github.composeweb/frontend/components/input"
	"io.github.composeweb/frontend/template"
)

const (
	usernameLabel = "Username or email"
	passwordLabel = "Password"
)

type loginDataStruct struct {
	Fields []input.InputData
	Error  string
	Success string
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	renderLoginPage(w, "", "", "")
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(usernameLabel)
	password := r.FormValue(passwordLabel)

	loginResponse, err := api.Login(username, password)
	if err != nil {
		renderLoginPage(w, username, err.Error(), "")
	}

	accessTokenCookie := http.Cookie{
		Name: "Authorization",
		Value: loginResponse.TokenType + " " + loginResponse.AccessToken,
		HttpOnly: true,
		Secure:   true,
		MaxAge:  loginResponse.ExpiresIn,
		Path: 	  "/",
		SameSite: http.SameSiteLaxMode,
	}
	expiresCookie := http.Cookie{
		Name: "expires",
		Value: fmt.Sprintf("%d", time.Now().Add(time.Second * time.Duration(loginResponse.ExpiresIn)).Unix()),
		HttpOnly: true,
		Secure:   true,
		Path: 	  "/",
		MaxAge:  loginResponse.ExpiresIn,
		SameSite: http.SameSiteLaxMode,
	}
	refreshTokenCookie := http.Cookie{
		Name: "refresh_token",
		Value: loginResponse.RefreshToken,
		HttpOnly: true,
		Secure:   true,
		Path: 	  "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &refreshTokenCookie)
	http.SetCookie(w, &accessTokenCookie)
	http.SetCookie(w, &expiresCookie)
	w.Header().Add("Authrorization", loginResponse.TokenType + " " + loginResponse.AccessToken)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func renderLoginPage(
	w http.ResponseWriter,
	username,
	error,
	success string,
) {
	data := loginDataStruct{
		Fields: []input.InputData{
			{
				Label:      usernameLabel,
				Type:       "text",
				Value:      username,
				IsRequired: true,
			},
			{
				Label:      passwordLabel,
				Type:       "password",
				Value:      "",
				IsRequired: true,
				Pattern: api.PASSWORD_REGEX,
				ValidatorHints: []string{
					"Invalid password format",
				},
			},
		},
		Error:  error,
		Success: success,
	}

	template.Render(w, "login.html", data)
}
