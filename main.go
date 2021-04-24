package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/tkanos/gonfig"

	"environments"
	"models"
)

var configuration environments.Configuration

func init() {
	err := gonfig.GetConf("./environments/production.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	usdPtr := flag.String("usd", "", "USD value to convert to EUR")
	flag.Parse()

	if *usdPtr == "" {
		log.Fatal("Missing -usd parameter with the amount of USD to convert")
	}

	usdAmount, _ := strconv.ParseFloat(*usdPtr, 64)
	latestUSDRate := getLatestedUSDRate()

	eurConvertedAmount := convertUSDtoEUR(usdAmount, latestUSDRate.Rate)

	fmt.Println(usdAmount, "USD = ", eurConvertedAmount, "EUR")
}

func getLatestedUSDRate() models.ExchangeratesapiResponse {
	endpoint := "latest"
	urlParams := "base=EUR&symbols=USD"
	url := buildUrlFor(endpoint, urlParams)

	fmt.Println("📡 Calling Exchangerateapi to get latest rates...")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Println("✅ Latest rates received: 1 EUR =", exchangeratesapiResponseObject.Rate, "USD")

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

func convertUSDtoEUR(usdAmount float64, rate float64) float64 {
	return usdAmount * rate
}
