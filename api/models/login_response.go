package models

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken string `json:"id_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}
