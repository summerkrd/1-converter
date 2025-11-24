package api_test

import (
	"2-calc/api"
	"2-calc/config"
	"encoding/json"
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
	testData := []byte(`{"test": "value", "number": 123}`)
	testName := "test-create-bin"

	client, err := CreateNewClient()
	if err != nil {
		t.Fatal(err.Error())
	}

	binID, err := client.CreateBin(testData, testName)
	if err != nil {
		t.Error(err.Error())
	}

	if binID == "" {
		t.Fatal("CreateBin вернул пустой ID")
	}

	defer func() {
		err := client.DeleteBin(binID)
		if err != nil {
			t.Errorf("не удалось удалить bin: %v", err)
		}
	}()

	binData, err := client.GetBin(binID)
	if err != nil {
		t.Error(err.Error())
	}
	if binData == nil {
		t.Error("binData is nil")
	}

	var unmTestData map[string]string
	err = json.Unmarshal(testData, &unmTestData)
	if binData["test"] != unmTestData["test"] {
		t.Error("ошибка: поле test не совпадает")
	}
	if binData["number"] != unmTestData["number"] {
		t.Error("ошибка: поле number не совпадает")
	}
}

func TestCreateBin(t *testing.T) {
	testData := []byte(`{"test": "value", "number": 123}`)
	testName := "test-create-bin"

	client, err := CreateNewClient()
	if err != nil {
		t.Fatal(err.Error())
	}

	binID, err := client.CreateBin(testData, testName)

	if err != nil {
		t.Error(err.Error())
	}

	if binID == "" {
		t.Fatal("CreateBin вернул пустой ID")
	}

	defer func() {
		err := client.DeleteBin(binID)
		if err != nil {
			t.Errorf("не удалось удалить bin: %v", err)
		}
	}()
}
