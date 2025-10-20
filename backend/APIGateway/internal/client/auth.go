package client

import (
	"apigateway/internal/config"
	"apigateway/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(cfg *config.Config) *AuthClient {
	return &AuthClient{
		baseURL: cfg.AuthAddress,
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
func (a *AuthClient) Register(email, password string) error {
	url := fmt.Sprintf("%s/api/register", a.baseURL)
	requestBody := models.AuthRequest{
		Email:    email,
		Password: password,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		errorMsg := string(body)
		return errors.New(errorMsg)
	}
	return nil
}

// func (a *AuthClient) Login(email, password string) (string, error) {
// 	url := fmt.Sprintf("%s/api/login")
// }
