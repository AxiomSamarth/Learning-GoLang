package server

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Login ...
func Login(host, username, password string) (statusCode int) {

	payload := make(map[string]string)
	payload["username"] = username
	payload["password"] = password

	data, _ := json.Marshal(payload)

	request, err := http.NewRequest("POST", host+"/api/v1/login", bytes.NewBuffer(data))
	if err == nil {
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}

		client.CloseIdleConnections()
		defer response.Body.Close()

		return response.StatusCode
	}
	panic(err)
}
