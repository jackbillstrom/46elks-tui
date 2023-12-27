package elks

import "net/http"

// Client is a struct that holds the API credentials, base URL and an HTTP client
type Client struct {
	Username string
	Password string
	BaseURL  string
	Client   *http.Client
}
