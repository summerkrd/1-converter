package main

import (
	"fmt"
)

func main() {

	const (
		USD = "USD"
		EUR = "EUR"
		RUB = "RUB"
	)

	currencyMap := map[string]float64{
		USD: 1.0,
		EUR: 0.8,
		RUB: 81.9,
	}

	calculateCurrency(userInput(currencyMap[USD], currencyMap[EUR], currencyMap[RUB]))
}

func userInput(usd float64, eur float64, rub float64) (count float64, original float64, target float64) {

	const (
		USDId = 1
		EURId = 2
		RUBId = 3
	)

	var originalID int
	var targetID int

	for {
		fmt.Printf("Из какой валюты перевести?\n(%d-USD/%d-EUR/%d-RUB)", USDId, EURId, RUBId)
		_, err := fmt.Scan(&originalID)

		if err != nil || originalID < USDId || originalID > RUBId {

			fmt.Println("Incorrect data")
			continue

		} else {

			break
		}
	}

	for {
		fmt.Println("Введите количество валюты: ")

		_, err := fmt.Scan(&count)

		if err != nil || count <= 0 {
			fmt.Println("Incorrect data")
			continue

		} else {

			break
		}
	}

	for {
		fmt.Printf("В какую валюту перевести?\n(%d-USD/%d-EUR/%d-RUB)", USDId, EURId, RUBId)
		_, err := fmt.Scan(&targetID)

		if err != nil || targetID < USDId || targetID > RUBId || targetID == originalID {
			fmt.Println("Incorrect data")
			continue

		} else {

			break
		}
	}

	switch originalID {
	case 1:
		original = usd
	case 2:
		original = eur
	case 3:
		original = rub
	}

	switch targetID {
	case 1:
		target = usd
	case 2:
		target = eur
	case 3:
		target = rub
	}

	return count, original, target
}

func calculateCurrency(count float64, original float64, target float64) {

	fmt.Print(count * (target / original))
}
