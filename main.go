package main

import "fmt"

func main() {

	const USD = 1
	const EUR = 2
	const RUB = 3

	const USDToEUR = 0.8515
	const USDToRUB = 81.9

	EURToRUB := USDToRUB / USDToEUR
	RubToEUR := USDToEUR / USDToRUB

	fmt.Println(EURToRUB)
	fmt.Println(RubToEUR)

}

func userInput() (count float64, original int, target int) {

	fmt.Println("Введите количество валюты: ")
	fmt.Scan(&count)

	fmt.Println("Из какой валюты перевести?\n(1-USD/2-EUR/3-RUB)")
	fmt.Scan(&original)

	fmt.Println("В какую валюту перевести?\n(1-USD/2-EUR/3-RUB)")
	fmt.Scan(&target)

	return
}

func calculateCurrency(count float64, original float64, target float64) {

}
