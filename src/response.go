package main

import (
	"io/ioutil"
	"net/http"
)

func getResponseBody(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}
