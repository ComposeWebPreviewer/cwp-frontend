package signup

import (
	"net/http"

	"io.github.composeweb/frontend/api"
	"io.github.composeweb/frontend/components/input"
	"io.github.composeweb/frontend/template"
)

const (
	emailLabel           = "Email"
	passwordLabel        = "Password"
	confirmPasswordLabel = "Confirm Password"
	usernameLabel        = "Username"
)

type signupPageData struct {
	Fields []input.InputData
	Error  string
	Success string
}

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	renderSignupPage(w, "", "", "", "", "")
}

func SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue(emailLabel)
	password := r.FormValue(passwordLabel)
	confirmPassword := r.FormValue(confirmPasswordLabel)
	username := r.FormValue(usernameLabel)

	if password != confirmPassword {
		renderSignupPage(
			w,
			email,
			"Passwords do not match",
			username,
			"",
			"",
		)
		return
	}

	_, err := api.Signup(email, password, username)
	if err != nil {
		renderSignupPage(
			w,
			email,
			"",
			username,
			err.Error(),
			"",
		)
		return
	}

	renderSignupPage(
		w,
		email,
		"",
		username,
		"",
		"Signup successful! You can now log in.",
	)
}

func renderSignupPage(
	w http.ResponseWriter,
	email,
	confirmPasswordError,
	username,
	error,
	success string,
) {
	data := signupPageData{
		Fields: []input.InputData{
			{
				Label:      emailLabel,
				Type:       "email",
				Value:      email,
				IsRequired: true,
			},
			{
				Label:      passwordLabel,
				Type:       "password",
				Value:      "",
				IsRequired: true,
				Pattern:    api.PASSWORD_REGEX,
				ValidatorHints: []string{
					"At least 8 characters",
					"At least one uppercase letter",
					"At least one lowercase letter",
					"At least one number",
				},
			},
			{
				Label:      confirmPasswordLabel,
				Type:       "password",
				Value:      "",
				IsRequired: true,
				Pattern:    api.PASSWORD_REGEX,
				ValidatorHints: []string{
					"Same as password",
				},
				Error:      confirmPasswordError,
			},
			{
				Label:      usernameLabel,
				Type:       "text",
				Value:      username,
				IsRequired: true,
				Pattern:    "^[a-zA-Z0-9_]{3,}$",
				ValidatorHints: []string{
					"At least 3 characters",
					"Only letters, numbers, and underscores",
				},
			},
		},
		Error: error,
		Success: success,
	}

	template.Render(w, "signup.html", data)
}
