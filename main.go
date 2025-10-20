package main

import (
	"fmt"
)

func main() {

	const (
		USD = 1
		EUR = 2
		RUB = 3
	)

	currencyMap := map[int]float64{
		USD: 1.0,
		EUR: 0.8,
		RUB: 81.9,
	}

	calculateCurrency(userInput(currencyMap))
}

func userInput(curr map[int]float64) (count float64, original float64, target float64) {

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

	return count, curr[originalID], curr[targetID]
}

func calculateCurrency(count float64, original float64, target float64) {

	fmt.Print(count * (target / original))
}
