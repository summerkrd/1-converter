package api

import (
	"2-calc/config"
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) CreateBin(data []byte, binName string) string {

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.apiKey)
	req.Header.Set("X-Bin-Name", binName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ошибка: " + err.Error())
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		fmt.Println("ошибка: статус код " + resp.Status)
		return ""
	}

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	var response Response
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		fmt.Println("ошибка: " + err.Error())
		return ""
	}

	return response.Metadata.ID
}

func (c *Client) GetBin(id string) {
	currentURL, err := url.Parse(c.baseURL + "/" + id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req, err := http.NewRequest("GET", currentURL.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("X-Master-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("StatusCode not 200")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ошибка: не удалось прочитать тело ответа")
		return
	}

	fmt.Println(string(body))
}

func (c *Client) UpdateBin(data []byte, binId string) {
	req, err := http.NewRequest("PUT", c.baseURL+"/"+binId, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ошибка: " + err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("ошибка: статус код " + resp.Status)
		return
	}
}

func (c *Client) DeleteBin() {

}
