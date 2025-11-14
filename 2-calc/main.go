package main

import (
	"2-calc/api"
	"2-calc/config"
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

}
