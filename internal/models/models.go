package models

// OutgoingSMS is a struct that holds the data for an outgoing SMS
type OutgoingSMS struct {
	From          string  `json:"from"`
	To            string  `json:"to"`
	Text          string  `json:"message"`
	DryRun        *string `json:"dryrun,omitempty"`        // Optional: "yes" to simulate sending
	WhenDelivered *string `json:"whendelivered,omitempty"` // Optional: URL for callback when delivered
	FlashSMS      *string `json:"flashsms,omitempty"`      // Optional: "yes" to send as flash SMS
	DontLog       *string `json:"dontlog,omitempty"`       // Optional: "message" för att inte logga meddelandetexten
}
