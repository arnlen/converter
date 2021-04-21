package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/tkanos/gonfig"

	"models"
	"environments"
)

var configuration environments.Configuration

func init() {
	err := gonfig.GetConf("./environments/production.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(getLatestedUSDRate())
}

func getLatestedUSDRate() models.ExchangeratesapiResponse {
	endpoint := "latest"
	urlParams := "base=EUR&symbols=USD"
	url := buildUrlFor(endpoint, urlParams)

	fmt.Println("Calling:", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:", string(responseData))

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

func buildUrlFor(endpoint string, urlParams string) string {
	baseApiUrl := configuration.ExchangeratesapiRoot
	accessKey := configuration.AccessKey
	integratedUrlParams := "?" + "access_key=" + accessKey

	if urlParams != "" {
		integratedUrlParams += "&" + urlParams
	}

	url := baseApiUrl + endpoint + integratedUrlParams

	return url
}