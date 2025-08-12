package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"io.github.composeweb/frontend/api/models"
)

var (
	signupEndpoint = fmt.Sprintf("%s/auth/signup", API_BASE_URL)
	loginEndpoint  = fmt.Sprintf("%s/auth/login", API_BASE_URL)
)

func Signup(email, password, username string) (models.SignupResponse, error) {
	requestBody := map[string]string{
		"email":    email,
		"password": password,
		"username": username,
	}
	jsonBody, _ := json.Marshal(requestBody)

	res, _ := http.Post(signupEndpoint, "application/json", bytes.NewBuffer(jsonBody))
	defer res.Body.Close()

	responseBody, _ := io.ReadAll(res.Body)

	var baseResponse models.BaseResponse
	json.Unmarshal(responseBody, &baseResponse)

	switch baseResponse.Status {
	case http.StatusOK:
	  var signupResponse models.SignupResponse
		json.Unmarshal(responseBody, &signupResponse)
		return signupResponse, nil

	case http.StatusConflict:
	  return models.SignupResponse{}, fmt.Errorf("Email or Username already exists, if you already have an account, please login")

	default:
		return models.SignupResponse{}, fmt.Errorf("Unexpected error")
	}
}

func Login(username, password string) (models.LoginResponse, error) {
	requestBody := map[string]string{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(requestBody)

	res, _ := http.Post(loginEndpoint, "application/json", bytes.NewBuffer(jsonBody))
	defer res.Body.Close()

	responseBody, _ := io.ReadAll(res.Body)

	var baseResponse models.BaseResponse
	json.Unmarshal(responseBody, &baseResponse)

	switch baseResponse.Status {
	case http.StatusOK:
	  var loginResponse models.LoginResponse
		json.Unmarshal(responseBody, &loginResponse)
		return loginResponse, nil
	
	case http.StatusNotFound:
		return models.LoginResponse{}, fmt.Errorf("Invalid username or password")

	case http.StatusForbidden:
		return models.LoginResponse{}, fmt.Errorf("Account is not activated, please check your email for the activation link")

	default:
		return models.LoginResponse{}, fmt.Errorf("Unexpected error")
	}
}
