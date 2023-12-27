package elks

import (
	"os"
	"testing"
)

// TestNewClient tests the NewClient function
func TestNewClient(t *testing.T) {
	username := "testuser"
	password := "testpassword"

	client := NewClient(username, password)

	if client.Username != username || client.Password != password {
		t.Errorf("NewClient did not set credentials correctly")
	}
	if client.BaseURL != "https://api.46elks.com/a1" {
		t.Errorf("NewClient did not set BaseURL correctly")
	}
	if client.Client == nil {
		t.Errorf("NewClient did not initialize http.Client")
	}
}

// TestCreateClient tests the CreateClient function
func TestCreateClient(t *testing.T) {
	expectedUsername := "test_api_user"
	expectedPassword := "test_api_pass"

	os.Setenv("46ELKS_API_USERNAME", expectedUsername)
	os.Setenv("46ELKS_API_PASSWORD", expectedPassword)
	defer func() {
		os.Unsetenv("46ELKS_API_USERNAME")
		os.Unsetenv("46ELKS_API_PASSWORD")
	}()

	client := CreateClient()

	if client.Username != expectedUsername || client.Password != expectedPassword {
		t.Errorf("CreateClient did not set credentials from environment variables correctly")
	}
}
