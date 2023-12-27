package ui

import (
	"46elks-tui/internal/models"
	"46elks-tui/pkg/elks"
)

// processOptions processes the options given to the program, such as dryrun and flashsms options when manually sending an SMS via the TUI
func processOptions(sms *models.OutgoingSMS, options []string) {
	for _, option := range options {
		switch option {
		case "dryrun":
			dryrun := "yes"
			sms.DryRun = &dryrun
		case "flashsms":
			flashsms := "yes"
			sms.FlashSMS = &flashsms
		}
	}
}

// sendSMS sends an SMS using the 46elks API, and returns the result as a string and an error if any
func sendSMS(client *elks.Client, sms *models.OutgoingSMS) (string, error) {
	err := client.SendSMS(sms)
	if err != nil {
		return "", err
	}
	return "SMS was sent successfully!", nil
}
