package utils

import (
	"46elks-tui/internal/models"
	"testing"
)

// TestValidateFrom tests the validateFrom function
func TestValidateFrom(t *testing.T) {
	tests := []struct {
		from   string
		expect bool
	}{
		{"", false},               // Testar tomt värde
		{"Sender", true},          // Testar giltigt värde
		{"Invalid Sender", false}, // Testar värde med mellanrum
	}

	for _, test := range tests {
		err := validateFrom(test.from)
		if (err == nil) != test.expect {
			t.Errorf("validateFrom(%s) = %v, want %v", test.from, err, test.expect)
		}
	}
}

// TestValidateTo tests the validateTo function
func TestValidateTo(t *testing.T) {
	tests := []struct {
		to     string
		expect bool
	}{
		{"", false},
		{"+46700000000", true},
		{"123456", false},
	}

	for _, test := range tests {
		err := validateTo(test.to)
		if (err == nil) != test.expect {
			t.Errorf("validateTo(%s) = %v, want %v", test.to, err, test.expect)
		}
	}
}

// TestValidateMessage tests the validateMessage function
func TestValidateMessage(t *testing.T) {
	tests := []struct {
		message string
		expect  bool
	}{
		{"Hello", false},
		{"Hello, this is a longer message.", true},
	}

	for _, test := range tests {
		err := validateMessage(test.message)
		if (err == nil) != test.expect {
			t.Errorf("validateMessage(%s) = %v, want %v", test.message, err, test.expect)
		}
	}
}

// TestValidateSMS tests the ValidateSMS function
func TestValidateSMS(t *testing.T) {
	tests := []struct {
		sms    models.OutgoingSMS
		expect bool
	}{
		{models.OutgoingSMS{From: "", To: "+46700000000", Text: "This is a test message"}, false},
		{models.OutgoingSMS{From: "Sender", To: "+46700000000", Text: "This is a test message"}, true},
		// Lägg till fler testfall här
	}

	for _, test := range tests {
		err := ValidateSMS(&test.sms)
		if (err == nil) != test.expect {
			t.Errorf("ValidateSMS() = %v, want %v", err, test.expect)
		}
	}
}
