package main

import (
    "log"
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
    "fmt"
    "os"
    "strings"
    "github.com/slack-go/slack"
)


const (
    // action is used for slack attament action.
    actionApprove = "approve"
    actionReject  = "reject"
)

type SlackListener struct {
	client    *slack.Client
}

// interactionHandler handles interactive message response.
type interactionHandler struct {
	slackClient       *slack.Client
}



func main() {
        // Listening slack event and response
        log.Printf("[INFO] Start slack event listening")

	token := os.Getenv("SLACKTOKEN")
	client := slack.New(token)

        slackListener := &SlackListener{
		client:    client,
	}
	go slackListener.ListenAndResponse()

        http.Handle("/", interactionHandler{})
        log.Printf("[INFO] Server listening on :%s", 8080)
        log.Fatal(http.ListenAndServe(":8080", nil))
}




// LstenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *SlackListener) ListenAndResponse() {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}




// handleMesageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
        rtm := s.client.NewRTM()
	var response string
	text := ev.Msg.Text
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	acceptedGreetings := map[string]bool{
		"what's up?": true,
		"hey!":       true,
		"yo":         true,
	}
	acceptedHowAreYou := map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
	}
        buildMessage := map[string]bool{
                "build": true,
        }

	// Parse message
	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, ev.Msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, ev.Msg.Channel))
	} else if buildMessage[text] {
	        attachment := slack.Attachment{
                   Color: "#36a64f",
                   CallbackID: "devdeploy",
		   Text:    "Deploying To Production Environment for `V23.456.7`",
                   Actions: []slack.AttachmentAction{
                     {
                       Name: actionApprove,
                       Text: "Approve",
                       Type: "button",
                       Value: "approve",
                       Style: "primary",
                     },
                     {
                       Name: actionReject,
                       Text: "Reject",
                       Type: "button",
                       Value: "reject",
                       Style: "danger",
                     },
                   },
                }
	        rtm.PostMessage(ev.Channel, slack.MsgOptionAttachments(attachment))
        }
    return nil
}




//interactionHandler handles message events.
func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        slackcommand := r.FormValue("command")

        if slackcommand == "/interaction" {
	   fmt.Fprint(w,"What's up buddy!?!?!")
           return
        } else if slackcommand == "/build" {
           fmt.Fprint(w,"Getting Ready to Build")
           return
        } else if slackcommand == "/deploy" {
	   fmt.Fprint(w,"Getting Ready to Deploy")
           return
        } else {
           fmt.Fprint(w,"I do not understand your command.")
           return
        }

        if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

        buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	action := message.Actions
        fmt.Println(action)
	switch action.Name {
	case actionApprove:
		title := ":ok: your order was submitted! yay!"
		responseMessage(w, message.OriginalMessage, title, "")
		return
	case actionReject:
		title := fmt.Sprintf(":x: @%s canceled the request")
		responseMessage(w, message.OriginalMessage, title, "")
		return
	default:
		log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// responseMessage response to the original slackbutton enabled message.
// It removes button and replace it with message which indicate how bot will work
func responseMessage(w http.ResponseWriter, original slack.Message, title, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&original)
}
