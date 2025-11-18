package main

import (
	"2-calc/api"
	"2-calc/config"
	"2-calc/data"
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
	fName := flag.String("name", "", "имя бина")
	fID := flag.Int("id", 0, "id")
	flag.Parse()

	LocalBinsData, err := data.ReadFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *fCreate != "" {
		if *fFile != "" && *fName != "" {
			userData, err := data.ReadUserJSON(*fFile)
			if err != nil {
				fmt.Println("ошибка: не удалось прочитать UserJSON")
				return
			}
			id := client.CreateBin(*userData, *fName)

			newBin := data.BinInfo{
				ID:   id,
				Name: *fName,
			}

			err = data.WriteFile(data.AddLocalBin(LocalBinsData, newBin))
			if err != nil {
				fmt.Println(err.Error())
			}
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
			client.GetBin(*fID)
		}

	} else if *fList != "" {
		for _, bin := range *LocalBinsData {
			fmt.Println("name: " + bin.Name + "\nid: " + bin.ID + "\n\n")
		}

	} else {
		fmt.Println("Введены не корректные данные")
		return
	}
}
