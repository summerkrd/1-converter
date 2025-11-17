package main

import (
	"2-calc/api"
	"2-calc/config"
	"flag"
	"fmt"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфига:", err)
		return
	}

	apiClient := api.NewClient(cfg)
	fmt.Println("API ключ загружен:", apiClient.GetKey())

	flagsHandler(apiClient)
}

func flagsHandler(client *api.Client) {
	fCreate := flag.String("create", "", "создать bin")
	fUpdate := flag.String("update", "", "обновить bin")
	fDelete := flag.String("delete", "", "удалить bin")
	fGet := flag.String("get", "", "получить bin")
	fList := flag.String("list", "", "список bins")
	fFile := flag.String("file", "", "файл")
	fName := flag.String("name", "", "имя файла")
	fID := flag.Int("id", 0, "id")
	flag.Parse()

	if *fCreate != "" {
		if *fFile != "" && *fName != "" {
			client.CreateBin()
		}
	} else if *fUpdate != "" {
		if *fFile != "" && *fID != 0 {
			client.UpdateBin()
		}
	} else if *fDelete != "" {
		if *fID != 0 {
			client.DeleteBin()
		}
	} else if *fGet != "" {
		if *fID != 0 {
			client.GetBin()
		}
	} else if *fList != "" {

	} else {
		fmt.Println("Введены не корректные данные")
		return
	}
}
