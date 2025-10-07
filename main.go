package main

import "fmt"

func main() {

	const USDToEUR = 0.8515
	const USDToRUB = 81.9

	EURToRUB := USDToRUB / USDToEUR

	fmt.Print(EURToRUB)

}

func userInput() (value float64) {

	fmt.Print("Введите количество валюты: ")

	fmt.Scan(&value)

	return
}

func calculateCurrency(count float64, original float64, target float64) {

}
