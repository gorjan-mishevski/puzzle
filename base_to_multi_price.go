package puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Currency - Normalized return data from BaseAgainsMultiPrice.
type Currency struct {
	base  string
	rates map[string]float64
}

// GetBase - Returns the base currency we are asking for rates. Ex: Base = "USD".
func (c *Currency) GetBase() string {
	return c.base
}

// GetRates - Returns the rates from the base currency to the requested ones. Ex: Base "USD" to "BTH, LIT"
func (c *Currency) GetRates() map[string]float64 {
	return c.rates
}

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
