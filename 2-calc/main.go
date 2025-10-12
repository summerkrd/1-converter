package main

import "fmt"

func main() {

	const AVGId = 1
	const SUMID = 2
	const MEDId = 3

	var currentId int

	var intSlice []int
	var strSlice []string

	for {

		fmt.Print("Выберите операцию: (1-Среднее/2-Сумма/3-Медиана)")

		_, err := fmt.Scan(&currentId)

		if currentId < AVGId || currentId > MEDId || err != nil {
			continue
		}

		break
	}

	switch currentId {

	case AVGId:

	case SUMID:

	case MEDId:

	}
}

func inputHandler(input string) {

}
