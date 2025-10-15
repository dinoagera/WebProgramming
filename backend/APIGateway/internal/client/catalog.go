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

type CatalogResponse struct {
	Status  string        `json:"status"`
	Catalog []models.Good `json:"catalog"`
}
type CatalogClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCatalogClient(cfg *config.Config) *CatalogClient {
	return &CatalogClient{
		baseURL: cfg.CatalogAddress,
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
func (c *CatalogClient) GetCatalog() ([]models.Good, error) {
	url := fmt.Sprintf("%s/api/getcatalog", c.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var catalogResp CatalogResponse
		if err := json.Unmarshal(body, &catalogResp); err != nil {
			return nil, err
		}
		return catalogResp.Catalog, nil

	case http.StatusInternalServerError:
		if string(body) == "Catalog is empty" {
			return []models.Good{}, nil
		}
		return nil, fmt.Errorf("catalog service error: %s", string(body))
	default:
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}
}
func (c *CatalogClient) GetImage(productID string) ([]byte, error) {
	if productID == "" {
		return nil, fmt.Errorf("product ID is required")
	}
	url := fmt.Sprintf("%s/api/image/%s", c.baseURL, productID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		contentType := resp.Header.Get("Content-Type")
		if contentType != "image/jpeg" {
			return nil, fmt.Errorf("unexpected content type: %s", contentType)
		}
		imageData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read image data: %w", err)
		}

		return imageData, nil

	case http.StatusBadRequest:
		return nil, fmt.Errorf("product ID is required")

	case http.StatusInternalServerError:
		body, _ := io.ReadAll(resp.Body)
		errorMsg := string(body)

		switch errorMsg {
		case "Image not found":
			return nil, fmt.Errorf("image not found for product %s", productID)
		default:
			return nil, fmt.Errorf("catalog service error: %s", errorMsg)
		}

	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}
}
