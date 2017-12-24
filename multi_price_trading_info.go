package puzzle

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// MultiPriceMarket - Containes the outermost layer of the response.
type MultiPriceMarket struct {
	// Type - This can be Ex: "RAW" or "DISPLAY"
	Type map[string]map[string]MultiPriceFrom
}

// MultiPriceFrom - Contains the name of the currency you take as base to compare to. Ex: "USD"
type MultiPriceFrom struct {
	From string
	// Contains list of the targetet currency.
	To map[string]MultiPriceTo
}

// MultiPriceTo - Contains the targetet currency. Ex: "ETH"
type MultiPriceTo struct {
	Name  string
	Field map[string]interface{}
}

// MultiPriceTradingInfo - Contains data from multiple currencies to multiple currencies.
func MultiPriceTradingInfo(from []string, to []string) MultiPriceMarket {
	url := fmt.Sprintf("%s/pricemultifull?fsyms=%s&tsyms=%s", baseHref, strings.Join(from, ","), strings.Join(to, ","))

	var unpacked map[string]map[string]map[string]map[string]interface{}
	body := makeRequest(url)
	json.Unmarshal(body, &unpacked)
	err := json.Unmarshal(body, &unpacked)
	if err != nil {
		log.Fatal("Failed unmarshaling data. \n", err)
	}

	mulPriceType := MultiPriceMarket{}
	sliceMulPriceCrypt := make(map[string]MultiPriceTo, 0)
	sliceMulPriceCurr := make(map[string]MultiPriceFrom, 0)
	mulPriceType.Type = make(map[string]map[string]MultiPriceFrom, 2)

	for from, to := range unpacked {
		for curr, crypt := range to {
			for name, val := range crypt {
				mulPriceCrypt := MultiPriceTo{
					Name:  name,
					Field: val,
				}

				sliceMulPriceCrypt[mulPriceCrypt.Name] = mulPriceCrypt
			}

			mulPriceCurr := MultiPriceFrom{
				From: curr,
				To:   sliceMulPriceCrypt,
			}

			sliceMulPriceCrypt = nil
			sliceMulPriceCrypt = make(map[string]MultiPriceTo, 0)

			sliceMulPriceCurr[mulPriceCurr.From] = mulPriceCurr
		}

		if mulPriceType.Type[from] == nil {
			mulPriceType.Type[from] = make(map[string]MultiPriceFrom, 1)
		}

		mulPriceType.Type[from] = sliceMulPriceCurr

		sliceMulPriceCurr = nil
		sliceMulPriceCurr = make(map[string]MultiPriceFrom, 0)

	}

	return mulPriceType
}
