package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"models"
)

func main() {
	fmt.Println(getLatestedUSDRate())
}

func getLatestedUSDRate() models.ExchangeratesapiResponse {
	access_key := os.Getenv("EXCHANGERATESAPI_ACCESS_KEY")
	base_api_url := "http://api.exchangeratesapi.io/v1/"
	endpoint_name := "latest"
	url_params := "?" + "access_key=" + access_key + "&base=EUR&symbols=USD"
	full_url := base_api_url + endpoint_name + url_params

	response, err := http.Get(full_url)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	// Unmarshal standard keys
	var exchangeratesapiResponseObject models.ExchangeratesapiResponse
	json.Unmarshal(responseData, &exchangeratesapiResponseObject)

	// Unmarshal nested rates
	var ratesObject map[string]interface{}
	json.Unmarshal(responseData, &ratesObject)

	rates := ratesObject["rates"].(map[string]interface{})

	// Should be a better way to achived this kind of association
	for key, value := range rates {
		exchangeratesapiResponseObject.Symbol = key
		exchangeratesapiResponseObject.Rate = value.(float64)
	}

	return exchangeratesapiResponseObject
}