package elks

import (
	"46elks-tui/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSendSMS tests the SendSMS method
func TestSendSMS(t *testing.T) {
	// Skapa en HTTP testserver
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Svara med status OK
		// Här kan du lägga till fler kontroller för inkommande förfrågan
	}))
	defer testServer.Close()

	// Skapa en klient med testserverns URL
	client := &Client{
		Username: "testuser",
		Password: "testpass",
		BaseURL:  testServer.URL + "/",
		Client:   &http.Client{},
	}

	// Skapa ett SMS-objekt för att testa
	testSMS := &models.OutgoingSMS{
		From: "Test",
		To:   "+46700000000",
		Text: "Detta är ett testmeddelande",
	}

	// Utför SendSMS-metoden
	err := client.SendSMS(testSMS)
	if err != nil {
		t.Errorf("SendSMS returnerade ett fel: %v", err)
	}
}
