package client

import "net/http"

type CartClient struct {
	baseURL    string
	httpClient *http.Client
}
