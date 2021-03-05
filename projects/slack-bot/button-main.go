package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/slack-go/slack"
)

const (
	actionApprove = "approve"
	actionReject  = "reject"
)


func main() {

	token := os.Getenv("SLACKTOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
        // Start listening slack events
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
                                respond(rtm, ev)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent) {
	var response string
	text := msg.Text
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


	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
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
	        rtm.PostMessage(msg.Channel, slack.MsgOptionAttachments(attachment))
        }


        /*action := msg.Attachments
        prettyJSON, err := json.MarshalIndent(action, "", "    ")
        if err != nil {
            log.Fatal("Failed to generate json", err)
        }
        fmt.Printf("%s\n", string(prettyJSON))*/
}
