package elks

import (
	"46elks-tui/internal/models"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SendSMS sends an SMS using the 46elks API
func (c *Client) SendSMS(sms *models.OutgoingSMS) error {
	data := url.Values{}
	data.Set("from", sms.From)
	data.Set("to", sms.To)
	data.Set("message", sms.Text)

	if sms.FlashSMS != nil {
		data.Set("flashsms", *sms.FlashSMS)
	}
	if sms.DryRun != nil {
		data.Set("dryrun", *sms.DryRun)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/sms", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send SMS: %s", string(body))
	}

	return nil
}
