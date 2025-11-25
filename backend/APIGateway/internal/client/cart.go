package client

import (
	"apigateway/internal/config"
	"apigateway/internal/models"
	"bytes"
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
		var cart models.CartResponse
		if err := json.Unmarshal(body, &cart); err != nil {
			return models.Cart{}, err
		}
		return cart.Cart, nil
	default:
		return models.Cart{}, err
	}
}
func (c *CartClient) AddItem(userID string, productID string, quantity int, price float64, category string) error {
	url := fmt.Sprintf("%s/api/additem", c.baseURL)
	requestBody := models.AddItemRequest{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		Category:  category,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("X-User-ID", userID)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", string(body))
	}
}
func (c *CartClient) RemoveItem(userID string, productID string) error {
	url := fmt.Sprintf("%s/api/removeitem", c.baseURL)
	requestBody := models.RemoveItemRequest{
		ProductID: productID,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
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
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", string(body))
	}
}
func (c *CartClient) UpdateItem(userID string, productID string, typeOperation int) error {
	url := fmt.Sprintf("%s/api/updateitem", c.baseURL)
	requestBody := models.UpdateItemRequest{
		ProductID:     productID,
		TypeOperation: typeOperation,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
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
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", string(body))
	}
}
func (c *CartClient) ClearCart(userID string) error {
	url := fmt.Sprintf("%s/api/clearcart", c.baseURL)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-User-ID", userID)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("clear cart failed: status %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}
