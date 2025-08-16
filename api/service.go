package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io.github.composeweb/frontend/api/models"
)

func GetCodespaces(authorization string) (models.CodespacesResponse, error) {
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, API_BASE_URL, nil)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var codespacesResponse models.CodespacesResponse
		if err := json.NewDecoder(resp.Body).Decode(&codespacesResponse); err != nil {
			return models.CodespacesResponse{}, fmt.Errorf("Failed to decode response")
		}
		return codespacesResponse, nil
	case http.StatusUnauthorized:
		return models.CodespacesResponse{}, fmt.Errorf("Unauthorized access, please login again")
	default:
		return models.CodespacesResponse{}, fmt.Errorf("Unexpected error occurred")
	}
}
