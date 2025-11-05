package api

import (
	"2-calc/config"
	"fmt"
)

type Client struct {
	apiKey string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		apiKey: cfg.Key,
	}
}

func (c *Client) GetKey() string {
	return c.apiKey
}

func (c *Client) SendRequest() {
	fmt.Println("Здесь будет HTTP запрос с ключом:", c.apiKey)
}
