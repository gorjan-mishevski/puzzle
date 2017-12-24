package puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// MultiPrice - Calculates the value for the searched currency. Return example: unpacked["ETH"]["USD"] gives the value ETH -> USD.
func MultiPrice(from []string, to []string) (unpacked map[string]map[string]float64) {
	url := fmt.Sprintf("%s/pricemulti?fsyms=%s&tsyms=%s", baseHref, strings.Join(from, ","), strings.Join(to, ","))

	body := makeRequest(url)
	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Fatal("Failed unmarshaling data. \n", err)
	}

	return
}
