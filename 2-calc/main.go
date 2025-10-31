package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	const AVGId = 1
	const SUMID = 2
	const MEDId = 3

	var currentId int

	var intSlice []int
	var strSlice string

	calculate := map[int]func([]int) string{
		1: calculateAvg,
		2: calculateSumm,
		3: calculateMed,
	}

	for {

		fmt.Print("Выберите операцию: (1-Среднее/2-Сумма/3-Медиана): ")

		_, err := fmt.Scan(&currentId)

		if currentId < AVGId || currentId > MEDId || err != nil {
			continue
		}

		break
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Введите числа через запятую:")

		scanner.Scan()
		strSlice = scanner.Text()

		var err error = nil
		intSlice, err = inputNumbersHandler(strSlice)

		if err != nil {
			continue
		}

		break
	}

	fmt.Print(calculate[currentId](intSlice))

	//switch currentId {
	//
	//case AVGId:
	//	fmt.Print("Среднее значение: ")
	//	fmt.Print(calculateAvg(intSlice))
	//
	//case SUMID:
	//	fmt.Print("Сумма: ")
	//	fmt.Print(calculateSumm(intSlice))
	//
	//case MEDId:
	//	fmt.Print("Медиана: ")
	//	fmt.Print(calculateMed(intSlice))
	//}
}

func inputNumbersHandler(input string) ([]int, error) {

	parts := strings.Split(input, ",")
	handled := make([]int, 0, len(parts))

	for i, v := range parts {

		parts[i] = strings.TrimSpace(v)
	}

	for _, v := range parts {

		n, err := strconv.Atoi(v)

		if err == nil {

			handled = append(handled, n)

		} else {

			err = errors.New("не корректный ввод")
			return nil, err
		}
	}

	return handled, nil
}

func calculateSumm(arr []int) string {

	sum := 0

	for _, v := range arr {
		sum += v
	}

	return fmt.Sprintln("Сумма: ", sum)
}

func calculateAvg(arr []int) string {

	sum := 0

	for _, v := range arr {
		sum += v
	}

	avg := float64(sum) / float64(len(arr))
	return fmt.Sprintln("Среднее значение: ", avg)
}

func calculateMed(arr []int) string {

	sort.Ints(arr)
	middle := len(arr) / 2

	if len(arr)%2 == 0 {
		return fmt.Sprintln("Медиана: ", float64(arr[middle-1]+arr[middle])/2.0)
	} else {
		return fmt.Sprintln("Медиана: ", float64(arr[middle]))
	}
}
