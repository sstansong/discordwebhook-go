
# discord-go-webhook

Go Lang library to send messages to Discord via Incoming Webhooks.

This package uses the framework from the slack-go-webhook package:
https://github.com/ashwanthkumar/slack-go-webhook

## Usage
```go
package main
import (
	"fmt"
	"github.com/sstansong/discord-go-webhook"
)

func main() {
	webhookUrl := "https://discordapp.com/api/webhooks/token/ID"
	payload := dcwebhook.Payload{
		Username:"root",
		TTS:false,
		IconUrl:"https://github.io",
		Content:"Hello from <github.com/sstansong/discord-go-webhook>, a Go-Lang library to send discord webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>",
		Embeds: []dcwebhook.Embed{
			{
				Title:"",
				URL:"",
				Description:"",
				Color:0,
				Author:,
				Fields:[]dcwebhook.Field{},
				Image:,
				Thumbnail:,
				Footer:,
			},
		},
	}

	err := dcwebhook.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
```