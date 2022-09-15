package main

import (
	"fmt"
	"src/currency"
)

func main() {
	symbol := [...]string{currency.USD: "$", currency.EUR: "9",
		currency.GBP: "!", currency.RMB: "¥"}
	fmt.Println(symbol[currency.USD])
	fmt.Println("value of currency.USD:", currency.USD)
}
