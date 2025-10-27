package storage

import (
	"3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBinToJSON(bin bins.Bin) {

	file, err := os.Create("bin.json")
	if err != nil {
		fmt.Println("Не удалось создать файл")
	}

	defer file.Close()

	data, err := json.Marshal(bin)
	if err != nil {
		fmt.Println("Не удалось преобразовать")
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Не удалось записать")
	}
}

func ReadBinFromJSON(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Не удалось прочитать")
	}

	var readingBin bins.Bin

	err = json.Unmarshal(data, &readingBin)
	if err != nil {
		fmt.Println("Не удалось преобразовать")
	}
}
