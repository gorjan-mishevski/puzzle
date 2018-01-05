package puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// PriceHistoryByDay - Given a base currency it will load the corresponding rate value based on that day.
// The day is calculated from the timeStamp and will provide the value for that day.
func PriceHistoryByDay(base string, to []string, timeStamp int32) *Currency {
	url := fmt.Sprintf("%s/pricehistorical?fsym=%s&tsyms=%s&ts=%v", baseHref, base, strings.Join(to, ","), timeStamp)
	body := makeRequest(url)

	var unpacked interface{}
	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Fatal("Failed unmarshaling data. \n", err)
	}

	unkRates := unpacked.(map[string]interface{})[base].(map[string]interface{})

	sanitizedCurrency := &Currency{
		base:  base,
		rates: make(map[string]float64),
	}

	for name, value := range unkRates {
		sanitizedCurrency.rates[name] = value.(float64)
	}

	return sanitizedCurrency
}
