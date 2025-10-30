package storage

import (
	"3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

type Bin bins.Bin

type Storage struct{}

type BinStorage interface {
	SaveBinToJSON(bin Bin) error
	ReadBinFromJSON(fileName string) (*Bin, error)
}

func (s *Storage) SaveBinToJSON(bin Bin) error {

	file, err := os.Create("bin.json")
	if err != nil {
		fmt.Println("Не удалось создать файл")
	}

	defer file.Close()

	data, err := json.Marshal(bin)
	if err != nil {
		fmt.Println("Не удалось преобразовать")
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Не удалось записать")
		return err
	}
	return nil
}

func (s *Storage) ReadBinFromJSON(fileName string) (*Bin, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Не удалось прочитать")
		return nil, err
	}

	var readingBin Bin

	err = json.Unmarshal(data, &readingBin)
	if err != nil {
		fmt.Println("Не удалось преобразовать")
		return nil, err
	}

	return &readingBin, nil
}
