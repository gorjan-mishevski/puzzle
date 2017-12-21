package puzzle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Get the price.
// {
// 	"ETH": {
// 		"USD": 736.83,
// 		"EUR": 619.8
// 	},
// 	"BTC": {
// 		"USD": 18565.05,
// 		"EUR": 15757.92
// 	},
// 	"REP": {
// 		"USD": 47.1,
// 		"EUR": 39.39
// 	}
// }
// https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH,BTC,REP&tsyms=USD,EUR

var (
	baseHref = "https://min-api.cryptocompare.com/data"
)

// MultiPrice - Calculates the value for the searched currency. Return example: unpacked["ETH"]["USD"] gives the value ETH -> USD.
func MultiPrice(from []string, to []string) (unpacked map[string]map[string]float64) {
	url := fmt.Sprintf("%s/pricemulti?fsyms=%s&tsyms=%s", baseHref, strings.Join(from, ","), strings.Join(to, ","))

	client := http.Client{
		Timeout: time.Second * 5, // 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Can't create new request for multi price. \n", err)
	}

	req.Header.Add("User-Agent", "web-app")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed feching from API. \n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed reading body. \n", err)
	}

	err = json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Fatal("Failed unmarshaling data. \n", err)
	}

	return
}

type Currency struct {
	base  string
	rates map[string]float64
}

func (c *Currency) GetBase() string {
	return c.base
}

func (c *Currency) GetRates() map[string]float64 {
	return c.rates
}

// BaseAgainsMultiPrice - Given a base currency it will load the corresponding rate value. Ex: "USD"(1$) -> "ETH"(0.00X), "BTC(0.00X)"
func BaseAgainsMultiPrice(base string, to []string) *Currency {
	url := fmt.Sprintf("%s/pricemulti?fsyms=%s&tsyms=%s", baseHref, base, strings.Join(to, ","))

	client := http.Client{
		Timeout: time.Second * 5, // 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Can't create new request for multi price. \n", err)
	}

	req.Header.Add("User-Agent", "web-app")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed feching from API. \n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed reading body. \n", err)
	}

	var unpacked interface{}
	err = json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Fatal("Failed unmarshaling data. \n", err)
	}

	sanitizedCurrency := &Currency{
		base:  base,
		rates: make(map[string]float64),
	}

	unkRates := (unpacked.(map[string]interface{})[base]).(map[string]interface{})
	for name, val := range unkRates {
		sanitizedCurrency.rates[name] = val.(float64)
	}

	return sanitizedCurrency
}
