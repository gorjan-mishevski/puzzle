# PUZZLE
Golang CryptoCompare library

# Why
If you are searching for a nice wrapper library that used CryptoCompare API this is the one. ☝️ 
I was not satisfied with how dynamic and unpredicted the API response was plus Go being a statically typed language you waste time configuring and normalizing data before you can actually use it.

This library has a lot of struct's (or type asserted response instead of interface{}) that extend the response, thus ensuring nice clean access to the data and providing easy way to add more functionalities with methods on the structs.

Currently I deliberatly don't add any interfaces and methods becouse I want you to use whatever methods fits best.

# Installation
```go
go get github.com/gorjan-mishevski/puzzle
```

# Usage
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
Let's see how to get from multiple currencies to multiples rates.<br>

```go
response := puzzle.MultiPrice([]string{"USD", "EUR"}, []string{"ETH", "BTC"})
```
1 Argument wants a slice of currencies to be taken as base.<br>
2 Argument is a slice of currencies you want the rate.<br>

Response:
```
response <map[string]map[string]float64>
        "USD": map[string]float64
                "ETH": 0.0000034
                "BTC": 0.0000000344
        "EUR": map[string]float64
                "ETH": 0.0000034
                "BTC": 0.0000000344
```


