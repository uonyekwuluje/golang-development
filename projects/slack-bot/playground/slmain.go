package main

import (
	"fmt"
	"os"
        "strings"
	"github.com/slack-go/slack"
)


func main() {

	token := os.Getenv("SLACKTOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
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

	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else {
                response = "Requested Command Not Available"
                rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
        }

        /*action := msg.Attachments
        prettyJSON, err := json.MarshalIndent(action, "", "    ")
        if err != nil {
            log.Fatal("Failed to generate json", err)
        }
        fmt.Printf("%s\n", string(prettyJSON))*/
}
