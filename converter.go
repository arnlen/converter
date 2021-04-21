package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	access_key := os.Getenv("EXCHANGERATESAPI_ACCESS_KEY")
	base_api_url := "http://api.exchangeratesapi.io/v1/"
	endpoint_name := "latest"
	url_params := "?" + "access_key=" + access_key + "&base=EUR&symbols=USD"
	full_url := base_api_url + endpoint_name + url_params

	response, err := http.Get(full_url)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}