package client

import (
	"apigateway/internal/config"
	"apigateway/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CartClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCartClient(cfg *config.Config) *CartClient {
	return &CartClient{
		baseURL: cfg.CartAddress,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     cfg.HTTPidleTimeout,
			},
		},
	}
}
func (c *CartClient) GetCart(userID string) (models.Cart, error) {
	url := fmt.Sprintf("%s/api/getcart", c.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Cart{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("X-User-ID", userID)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return models.Cart{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Cart{}, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var cart models.Cart
		if err := json.Unmarshal(body, &cart); err != nil {
			return models.Cart{}, err
		}
		return cart, nil
	default:
		return models.Cart{}, err
	}
}
