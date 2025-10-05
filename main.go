package main

import "fmt"

func main() {

	const USDToEUR = 0.8515
	const USDToRUB = 81.9

	EURToRUB := USDToRUB / USDToEUR

	fmt.Print(EURToRUB)
}
