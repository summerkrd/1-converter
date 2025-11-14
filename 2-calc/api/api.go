package api

import (
	"2-calc/config"
	"fmt"
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

func (c *Client) GetBin() {

}

func (c *Client) UpdateBin() {

}

func (c *Client) DeleteBin() {

}
