package utils

import (
	"io"
	"net/http"
)

// Makes an HTTP request to API Ninjas
func ApiNinjasRequest(api string) string {
	iniData := GetIniData()

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.api-ninjas.com/v1/"+api, nil)
	req.Header.Set("X-Api-Key", iniData.Section("api").Key("api_ninjas_key").String())
	res, _ := httpClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	return bodyString
}
