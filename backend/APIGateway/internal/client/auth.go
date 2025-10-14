package client

import "net/http"

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}
