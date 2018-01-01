
<p align="center">
        <a href="https://ibb.co/fH5Azw" target="_blank">
                <img width="350" src="https://preview.ibb.co/ejVcew/Puzzle_Crypto_Logo.png" alt="Puzzle_Crypto_Logo" border="0">
        </a>
</p>

<p align="center">
        <a href="https://goreportcard.com/report/github.com/gorjan-mishevski/puzzle">
                <img src="https://goreportcard.com/badge/github.com/gorjan-mishevski/puzzle">
        </a>
         <a href="https://godoc.org/github.com/gorjan-mishevski/puzzle">
                <img src="https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square">
        </a>
        <a href="https://img.shields.io/badge/License-MIT-yellow.svg">
                <img src="https://opensource.org/licenses/MIT">
        </a>
</p>
     
# PUZZLE
Golang CryptoCompare library

If you are searching for a nice wrapper library that used CryptoCompare API this is the one. ☝️

This library has a lot of struct's (or type asserted response instead of interface{}) that extend the response, thus ensuring nice clean access to the data and providing easy way to add more functionalities with methods on the structs.

Currently I deliberatly don't add any interfaces and methods becouse I want you to use whatever methods fits best.

# Installation
```go
go get github.com/gorjan-mishevski/puzzle
```

# Usage

<h3> Example 1 </h3>
Let's start simple. Get the current value from a given currency.<br>
EX: I want to know what is the rate from US to ETH, BTC.<br><br>
<b>1 USD dollar</b> 💵 <b>-> ETH</b> 💎

```go
response := puzzle.BaseAgainsMultiPrice("USD", []string{"ETH", "BTC"})
```
1 Argument is the currency you take as base. <br>
2 Argument is a slice of currencies you want the rate.

Response:
```
response <*puzzle.Currency>
        base: "USD"
        rates: <map[string]float64>
                ["ETH"]: 0.001556
                ["BTC"]: 0.000002355
```
<hr>
<h3> Example 2 </h3>
Let's see how to get from multiple currencies to multiples rates.<br>

```go
response := puzzle.MultiPrice([]string{"USD", "EUR"}, []string{"ETH", "BTC"})
```
1 Argument wants a slice of currencies to be taken as base.<br>
2 Argument is a slice of currencies you want the rate.<br>

Response:
```
response <map[string]map[string]float64>
        "USD": <map[string]float64>
                "ETH": 0.0000034
                "BTC": 0.0000000344
        "EUR": <map[string]float64>
                "ETH": 0.0000034
                "BTC": 0.0000000344
```
<hr>
<h3> Example 3 </h3>
Now let's get more complicated. Let's say we want a complete trading information with dynamic to -> from currencies.
<br>

```go
response := puzzle.MultiPriceTradingInfo([]string{"USD", "EUR"}, []string{"ETH", "BTC"})
```

Response:<br>
```
response <puzzle.MultiPriceMarket>
        "RAW": <map[string]puzzle.MultiPriceFrom>
                "USD": <puzzle.MultiPriceFrom>
                        From: "USD"
                        To: <map[string]puzzle.MultiPriceTo>
                                "ETH": <puzzle.MultiPriceTo>
                                        Name: "ETH"
                                        Field: <map[string]interface{}>
                                                ... Here we have random field data.
                                                For now they are as interface but they can be
                                                mapped.
                                "BTC": <puzzle.MultiPriceTo>
                                        ...
                "EUR": <puzzle.MultiPriceFrom>
                        ...
                        
        "DYNAMIC": ...
```
<hr>
To be continued...

