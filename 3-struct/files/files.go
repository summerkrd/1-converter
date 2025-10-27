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

	var v any
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("Файл не является JSON")
	}
}
