package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "errors"
    "log"
    "os"
    "net/http"
    "time"
)


// SendSlackNotificationInput ...
type SendSlackNotificationInput struct {
	WebhookURL string
	Message    string
	Channel    string
	Username   string
	IconEmoji  string
}


// Attachment contains all the information for an attachment
type Attachment struct {
	Color    string `json:"color,omitempty"`
	Fallback string `json:"fallback"`
	CallbackID string `json:"callback_id,omitempty"`
	ID         int    `json:"id,omitempty"`
}



// SlackRequestBody ...
type SlackRequestBody struct {
	Text      string `json:"text"`
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
        Attachments []Attachment `json:"attachments,omitempty"`
}

func SendSlackNotification(input *SendSlackNotificationInput) error {
	slackBody, err := json.Marshal(&SlackRequestBody{
		Text:      input.Message,
		Channel:   input.Channel,
		Username:  input.Username,
		IconEmoji: input.IconEmoji,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, input.WebhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(buf) != "ok" {
		return errors.New("Not OK")
	}

	return nil
}



func main() {
    webhookURL := os.Getenv("SLACK_WEBHOOK")
    err := SendSlackNotification(&SendSlackNotificationInput{
		WebhookURL: webhookURL,
		Message:    "Hello world. Beeses",
		Channel:    "devops",
		Username:   "alice",
		IconEmoji:  "ghost",
    })

    if err != nil {
        log.Fatal(err)
    }
}
