package elks

import (
	"net/http"
	"os"
)

// NewClient returns a new Client struct
func NewClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
		BaseURL:  "https://api.46elks.com/a1",
		Client:   &http.Client{},
	}
}

// CreateClient returns a new Client struct
func CreateClient() *Client {
	apiUsername := os.Getenv("46ELKS_API_USERNAME")
	apiPassword := os.Getenv("46ELKS_API_PASSWORD")
	return NewClient(apiUsername, apiPassword)
}
