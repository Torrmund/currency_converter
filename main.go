package main

import (
	"fmt"
	"strings"
)

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"

	usdToEur = 0.87
	usdToRub = 78.50
)

func main() {
	eurToRub := usdToRub / usdToEur

	fmt.Println("====== КАЛЬКУЛЯТОР ВАЛЮТ ======")

	fmt.Printf("Доступные курсы валют:\n")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

	fmt.Println("\n" + strings.Repeat("=", 30))

	sourceCurrency := getCurrencyInput("исходную")
	amount := getAmountInput()
	targetCurrency := getCurrencyInput("целевую")

	if sourceCurrency == targetCurrency {
		fmt.Println("Ошибка: исходная и целевая валюта не могут совпадать!")
		return
	}

	result := convertCurrency(sourceCurrency, amount, targetCurrency)

	fmt.Println("\n" + strings.Repeat("=", 30))
	fmt.Println("РЕЗУЛЬТАТ КОНВЕРТАЦИИ:")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, sourceCurrency, result, targetCurrency)
	fmt.Println(strings.Repeat("=", 30))

}

func getCurrencyInput(currencyType string) string {
	var currency string

	for {
		fmt.Printf("\nВведите %s валюту (USD, EUR, RUB): ", currencyType)
		fmt.Scan(&currency)

		currency = strings.ToUpper(currency)

		if currency == USD || currency == EUR || currency == RUB {
			return currency
		}

		fmt.Printf("Ошибка: валюта '%s' не поддерживается.\n", currency)
		fmt.Println("Доступные валюты: USD, EUR, RUB")
	}
}

func getAmountInput() float64 {
	var amount float64

	for {
		fmt.Print("\nВведите сумму для конвертации (положительное число): ")
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка: введите корректное число!")
			continue
		}

		if amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть положительным числом!")
			continue
		}

		return amount
	}
}

func convertCurrency(sourceCurrency string, amount float64, targetCurrency string) float64 {
	var amountInUSD float64

	switch sourceCurrency {
	case USD:
		amountInUSD = amount
	case EUR:
		amountInUSD = amount / usdToEur
	case RUB:
		amountInUSD = amount / usdToRub
	}

	var result float64

	switch targetCurrency {
	case USD:
		result = amountInUSD
	case EUR:
		result = amountInUSD * usdToEur
	case RUB:
		result = amountInUSD * usdToRub
	}

	return result
}
