package client

import (
	"fmt"
	"net/http"
	"paymentservice/internal/config"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     cfg.IdleTimeout,
			},
		},
	}
}

func (c *Client) GetTotalPrice(userID string) (float64, error) {
	url := fmt.Sprintf("%s/api/")
}
