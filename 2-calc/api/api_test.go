package api_test

import (
	"2-calc/api"
	"2-calc/config"
	"errors"
	"testing"
)

func CreateNewClient() (*api.Client, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, errors.New("ошибка загрузки конфига")
	}

	apiClient := api.NewClient(cfg)
	return apiClient, nil
}

func TestGetBin(t *testing.T) {
	id := "12345"

	client, err := CreateNewClient()
	if err != nil {
		t.Error(err.Error())
	}

	client.GetBin(id)
}
