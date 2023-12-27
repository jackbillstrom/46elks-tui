package utils

import (
	"46elks-tui/internal/models"
	"errors"
	"flag"
	"strings"
	"unicode"
)

// ValidateSMS validates the SMS struct
func ValidateSMS(sms *models.OutgoingSMS) error {
	if err := validateFrom(sms.From); err != nil {
		return err
	}
	if err := validateTo(sms.To); err != nil {
		return err
	}
	if err := validateMessage(sms.Text); err != nil {
		return err
	}
	return nil
}

// ParseFlagsAndValidate parses the flags and validates the SMS struct
func ParseFlagsAndValidate() (*models.OutgoingSMS, error) {
	senderPtr := flag.String("from", "", "Sender ID or phone number")
	receiverPtr := flag.String("to", "", "Receiver phone number, in international format")
	textPtr := flag.String("text", "", "Content of the SMS")
	flashPtr := flag.Bool("flash", false, "Send as a flash SMS")
	dryrunPtr := flag.Bool("dryrun", false, "Dry run, don't send the SMS")
	flag.Parse()

	sms := &models.OutgoingSMS{
		From: *senderPtr,
		To:   *receiverPtr,
		Text: *textPtr,
	}

	if *flashPtr {
		flash := "yes"
		sms.FlashSMS = &flash
	}
	if *dryrunPtr {
		dryrun := "yes"
		sms.DryRun = &dryrun
	}

	if err := ValidateSMS(sms); err != nil {
		return nil, err
	}

	return sms, nil
}

// validateFrom checks that 'from' is not empty
func validateFrom(from string) error {
	if len(from) == 0 {
		return errors.New("sender ID must be set")
	}
	if strings.Contains(from, " ") {
		return errors.New("sender ID must not contain spaces")
	}
	return nil
}

// validateTo checks that 'to' is a valid international phone number
func validateTo(to string) error {
	if len(to) < 2 || to[0] != '+' || !isNumeric(to[1:]) {
		return errors.New("receiver must be a valid international phone number")
	}
	return nil
}

// validateMessage checks that 'message' is at least 10 characters long
func validateMessage(message string) error {
	if len(message) < 10 {
		return errors.New("message must be at least 10 characters long")
	}
	return nil
}

// isNumeric checks if a string only contains digits
func isNumeric(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
