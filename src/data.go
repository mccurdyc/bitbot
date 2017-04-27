package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func getHistoricalData(start, end string) []*gabs.Container {
	url := "http://api.coindesk.com/v1/bpi/historical/close.json"
	dateRangeUrl := fmt.Sprintf("%s?start=%s&end=%s", url, start, end)

	data := getResponseBody(dateRangeUrl)
	json, err := gabs.ParseJSON(data)

	if err != nil {
		panic(err)
	}

	bpi, err := json.Path("bpi").Children()
	if err != nil {
		panic(err)
	}

	return bpi
}

func getPriceOnDate(date string) float64 {
	url := "http://api.coindesk.com/v1/bpi/historical/close.json"
	dateUrl := fmt.Sprintf("%s?start=%s&end=%s", url, date, date)

	data := getResponseBody(dateUrl)
	json, err := gabs.ParseJSON(data)

	if err != nil {
		panic(err)
	}

	bpi, err := json.Path("bpi").Children()
	if err != nil {
		panic(err)
	}
	return bpi[0].Data().(float64)
}
