### Slack Bots
Codes and snippets for building slack bots

### Channel Message
To send a message to a channel, create your json pay load and send
```
SLACKFILE=""
SLACKWEBHOOK=""
curl -X POST -H 'Content-type: application/json' --data-binary "@${SLACKFILE}" ${SLACKWEBHOOK}
```

### Slack Buttons
This code snippet responds to message button above. Download ngrok and run in a separate terminal
```
ngrok http 8080
```
Then start slack bot
```
go run slack-button.go
```
