package data

import (
	"encoding/json"
	"errors"
	"os"
)

type BinInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const fileName = "bins.json"

func ReadFile() (*[]BinInfo, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("ошибка: не удалось прочитать json")
	}

	var bins []BinInfo
	err = json.Unmarshal(data, &bins)
	if err != nil {
		return nil, errors.New("ошибка: не удалось преобразовать из json")
	}

	return &bins, nil
}

func WriteFile(bins *[]BinInfo) error {
	newData, err := json.Marshal(bins)
	if err != nil {
		return errors.New("ошибка: не удалось преобразовать в json")
	}

	err = os.WriteFile(fileName, newData, 0644)
	if err != nil {
		return errors.New("ошибка: не удалось записать json")
	}
	return nil
}
