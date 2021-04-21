package main

import (
	"testing"
	"models"

	"github.com/jarcoal/httpmock"
)

func TestGetLatestedUSDRate(t *testing.T) {
	mockResponsePayload := `{
		"success":true,
		"timestamp":1619013243,
		"base":"EUR",
		"date":"2021-04-21",
		"rates":{
			"USD":2.50105
		}
	}`

	httpmock.Activate()

	httpmock.RegisterResponder("GET", `=~^http:\/\/api.exchangeratesapi.io\/v1\/latest.*`,
		httpmock.NewStringResponder(200, mockResponsePayload))

	expected := models.ExchangeratesapiResponse {
		Date:		"2021-04-21",
		Base:		"EUR",
		Symbol:	"USD",
		Rate: 	2.50105,
	}

	msg := getLatestedUSDRate()

	httpmock.DeactivateAndReset()

	if msg != expected {
		t.Fatalf(`GetLatestedUSDRate() = %v expected %v`, msg, expected)
	}
}