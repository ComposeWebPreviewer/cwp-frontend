package api

import "net/http"

func GetCodespaces(authorization string) {
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, API_BASE_URL, nil)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()


}
