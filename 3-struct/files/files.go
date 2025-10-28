package files

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) {

	if !isJSON(name) {
		fmt.Println("file is not JSON")
		return
	}

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Не удалось прочитать")
	}

	isValid, file := checkValidJSON(data)
	if !isValid {
		fmt.Println("Не валидный JSON")
		return
	}

	fmt.Println(file)
}

func checkValidJSON(data []byte) (isValid bool, v any) {

	err := json.Unmarshal(data, v)
	if err != nil {
		return false, nil
	}

	return
}

func isJSON(fileName string) bool {
	return strings.HasSuffix(fileName, ".json")
}
