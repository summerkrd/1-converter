package api

import (
	"2-calc/config"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey  string
	baseURL string
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

func (c *Client) SendRequest() {
	fmt.Println("Здесь будет HTTP запрос с ключом:", c.apiKey)
}

func (c *Client) CreateBin() {

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

func (c *Client) UpdateBin() {

}

func (c *Client) DeleteBin() {

}
