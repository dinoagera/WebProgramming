package client

import (
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

// func (c *Client) GetAmountPrice(userID string) (float64, error) {

// }
