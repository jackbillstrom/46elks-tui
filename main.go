package main

import (
	"46elks-tui/internal/models"
	"46elks-tui/internal/utils"
	"46elks-tui/pkg/elks"
	"fmt"
)

func main() {
	sms, err := utils.ParseFlagsAndValidate()
	if err != nil {
		fmt.Println("Error while validating flags:", err)
		displayTUI()
		return
	}

	client := elks.CreateClient()
	if sms.Text != "" {
		sendSMS(client, sms)
	} else {
		displayTUI()
	}
}

// sendSMS sends an SMS using the 46elks API, and prints the result to stdout
func sendSMS(client *elks.Client, sms *models.OutgoingSMS) {
	fmt.Println("Sending SMS...")
	err := client.SendSMS(sms)
	if err != nil {
		fmt.Println("Error sending SMS:", err)
		return
	}
	fmt.Println("SMS was sent successfully!")
}

func displayTUI() {
	fmt.Println("Showing TUI...")
	// TODO: Handle TUI, i.e. display a menu and send SMSes based on user input.
}
