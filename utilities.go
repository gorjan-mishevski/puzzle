package puzzle

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	baseHref = "https://min-api.cryptocompare.com/data"
)

func makeRequest(url string) []byte {
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

	return body
}
