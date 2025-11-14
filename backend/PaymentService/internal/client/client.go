package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"paymentservice/internal/config"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
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
	url := fmt.Sprintf("http://api_cart:8080/api/gettotalprice?user_id=%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	var response struct {
		Price float64 `json:"price"`
	}
	if resp.StatusCode != http.StatusOK {
		errorMsg := string(body)
		return 0.0, errors.New(errorMsg)
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return 0.0, err
	}
	return response.Price, nil
}
func (c *Client) ClearCart(userID string) error {
	url := "http://api_cart:8080/api/clearcart"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-ID", userID)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errorMsg := string(body)
		return errors.New(errorMsg)
	}
	return nil
}
