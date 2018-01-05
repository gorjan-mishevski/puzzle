package puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// BaseAgainsMultiPrice - Given a base currency it will load the corresponding rate value. Ex: "USD"(1$) -> "ETH"(0.00X), "BTC(0.00X)"
func BaseAgainsMultiPrice(base string, to []string) *Currency {
	url := fmt.Sprintf("%s/pricemulti?fsyms=%s&tsyms=%s", baseHref, base, strings.Join(to, ","))

	var unpacked interface{}
	body := makeRequest(url)
	err := json.Unmarshal(body, &unpacked)
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
