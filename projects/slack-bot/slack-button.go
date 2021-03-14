package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "github.com/slack-go/slack"
)


// event struct for slack event verification
type event struct {
        Type     string
        Text     string
        Channel  string
        ThreadTS string `json:"thread_ts"`
}

type eventMessage struct {
        Challenge   string
        Event       event
        AuthedUsers []string `json:"authed_users"`
}

type interactionHandler struct {
	verificationToken string
}

type buttonCapture []struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}


var (
    SLACK_BOT_TOKEN = "" // Paste your bot user token here
    CHANNEL_ID = "" // Paste your channel id here
    VERIFICATION_TOKEN = ""  //verification token
)


func main() {
    http.Handle("/receive", interactionHandler{
		verificationToken: VERIFICATION_TOKEN,
    })
    fmt.Println("[INFO] Server listening")
    http.ListenAndServe(":8080", nil)
}


func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        slackPayload := r.FormValue("payload")

        //fmt.Println(slackPayload)
        var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(slackPayload), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
	}

        // Convert actions to JSON
        data, err := json.Marshal(payload.ActionCallback)
        if err != nil {
            log.Fatal("Failed to generate json", err)
        }
        //fmt.Printf("%s\n", data)

        var buttonPayload buttonCapture
        json.Unmarshal(data, &buttonPayload)
        //fmt.Println("Button Value => ",buttonPayload[0].Value)
	fmt.Printf("Message button pressed by user %s with value %s\n", payload.User.Name, buttonPayload[0].Value)

        action := buttonPayload[0].Value
	switch action {
	case "approve":
		w.Write([]byte(":ok: Great. Release to Production"))
	case "reject":
		w.Write([]byte(":x: Deployment Cancled!"))
	default:
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("could not process callback: %s", action)))
		return
	}

}
