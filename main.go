package main

import "fmt"

const (
		usdToEur = 0.87
		usdToRub  = 78.50
	)

func main() {
	eurToRub := usdToRub / usdToEur

	fmt.Printf("Курсы валют:\n")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

	const amountEur = 500.0
	amountRub := amountEur * eurToRub
	fmt.Printf("\nПример конвертации:\n")
	fmt.Printf("%.2f EUR = %.2f RUB\n", amountEur, amountRub)
}

func getUserInput() (sourceCurrency string, amountCurrency float64, targetCurrency string) {
	fmt.Print("Введите название исходной валюты: ")
	fmt.Scan(&sourceCurrency)
	fmt.Print("Введите количество исходной валюты: ")
	fmt.Scan(&amountCurrency)
	fmt.Print("Введите название целевой валюты: ")
	fmt.Scan(&targetCurrency)
	return
}

func currencyConversion() {}
