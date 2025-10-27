package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) {

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Не удалось прочитать")
	}

	isJSON, _ := checkJSON(data)

	if !isJSON {
		fmt.Println("Файл не является JSON")
		return
	}
}

func checkJSON(data []byte) (isJSON bool, v any) {

	err := json.Unmarshal(data, v)
	if err != nil {
		return false, nil
	}
	return
}
