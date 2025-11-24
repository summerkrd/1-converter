package api

import (
	"2-calc/config"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey  string
	baseURL string
}

type Response struct {
	Metadata struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"metadata"`
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		apiKey:  cfg.Key,
		baseURL: "https://api.jsonbin.io/v3/b",
	}
}

func (c *Client) GetKey() string {
	return c.apiKey
}

func (c *Client) CreateBin(data []byte, binName string) (string, error) {

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(data))
	if err != nil {
		return "", errors.New("ошибка: не удалось отправить request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.apiKey)
	req.Header.Set("X-Bin-Name", binName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("ошибка: не удалось получить ответ от API")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return "", errors.New("ошибка: статус код: " + resp.Status)
	}

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("ошибка: не удалось прочитать resp.Body")
	}

	var response Response
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return "", errors.New("ошибка: не удалось преобразовать из json")
	}

	return response.Metadata.ID, nil
}

func (c *Client) GetBin(id string) (map[string]any, error) {
	currentURL, err := url.Parse(c.baseURL + "/" + id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", currentURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("StatusCode not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ошибка: не удалось прочитать тело ответа")
	}

	var binData map[string]any
	err = json.Unmarshal(body, &binData)
	if err != nil {
		return nil, errors.New("ошибка: не удалось преобразовать из json")
	}

	return binData, nil
}

func (c *Client) UpdateBin(data []byte, binId string) error {
	req, err := http.NewRequest("PUT", c.baseURL+"/"+binId, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("ошибка: " + err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("ошибка: статус код " + resp.Status)
	}
	return nil
}

func (c *Client) DeleteBin(id string) error {
	return nil
}
