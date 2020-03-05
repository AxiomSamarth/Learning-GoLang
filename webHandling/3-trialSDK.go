package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Authorization struct {
	Result string
}

func main() {
	fmt.Println("Let's make an API call to create ODIM session in GoLang!")
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	url := "https://localhost/api/v1/auth/token"
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("X-Auth-Username", "admin")
	req.Header.Set("X-Auth-Password", "plexxi")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,hi;q=0.8")

	checkError(err)

	fmt.Println("Create client")
	client := &http.Client{Transport: transCfg}

	fmt.Println("Send request")
	response, err := client.Do(req)
	checkError(err)

	var auth Authorization
	bytes, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bytes, &auth)
	checkError(err)
	fmt.Println(auth.Result)

	XAuthToken := string(bytes)
	fmt.Println("X-Auth-Token:", XAuthToken, response.Status)
	defer response.Body.Close()
	checkError(err)

	req, err = http.NewRequest("GET", "https://localhost/api/v1/integration_sets", nil)
	req.Header.Set("Authorization", XAuthToken)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,hi;q=0.8")

	response, err = client.Do(req)
	checkError(err)
	fmt.Println(response.Status)

	byteReponse, err := ioutil.ReadAll(response.Body)
	checkError(err)

	content := string(byteReponse)
	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
