package main

import (
	"2-calc/api"
	"2-calc/config"
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфига:", err)
		return
	}

	apiClient := api.NewClient(cfg)
	fmt.Println("API ключ загружен:", apiClient.GetKey())

	calculate := map[string]func([]int) string{
		"avg": calculateAvg,
		"sum": calculateSumm,
		"med": calculateMed,
	}

	var userInput string

	var intSlice []int
	var strSlice string

	for {

		fmt.Print("Выберите операцию: (avg - Среднее/sum - Сумма/med - Медиана): ")

		_, err := fmt.Scan(&userInput)

		if userInput == "" || err != nil {
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

	fmt.Print(calculate[userInput](intSlice))
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
